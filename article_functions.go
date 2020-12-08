package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM articles ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	article := Article{}
	res := []Article{}
	for selDB.Next() {
		var id int
		var title, content string
		err = selDB.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}
		article.Id = id
		article.Title = title
		article.Content = content
		res = append(res, article)
	}
	defer db.Close()
	json.NewEncoder(w).Encode(res)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	db := dbConn()
	selDB, err := db.Query("SELECT * FROM articles WHERE id = ?", key)
	if err != nil {
		panic(err.Error())
	}
	article := Article{}
	for selDB.Next() {
		var id int
		var title, content string
		err = selDB.Scan(&id, &title, &content)
		if err != nil {
			panic(err.Error())
		}
		article.Id = id
		article.Title = title
		article.Content = content
	}
	defer db.Close()
	json.NewEncoder(w).Encode(article)
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	a := Article{}
	json.NewDecoder(r.Body).Decode(&a)
	query, err := db.Prepare("INSERT INTO articles(title, content) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}
	query.Exec(a.Title, a.Content)
	defer db.Close()
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	a := Article{}
	json.NewDecoder(r.Body).Decode(&a)
	query, err := db.Prepare("UPDATE articles SET title = ?, content = ?, updated_at = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	updated_at := time.Now().Format("2006-01-02 15:04:05")
	_, err = query.Exec(
		a.Title,
		a.Content,
		updated_at,
		mux.Vars(r)["id"],
	)

	if err != nil {
		panic(err)
	}
	defer db.Close()
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	db := dbConn()

	selDB, err := db.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	selDB.Exec(key)
	defer db.Close()
}
