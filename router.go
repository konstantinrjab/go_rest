package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}