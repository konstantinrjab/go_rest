package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
}

func dbConn() (db *sql.DB) {
	dbUser := "user"
	dbPass := "password"
	dbName := "gotour"
	db, err := sql.Open("mysql", dbUser+":"+dbPass+"@tcp(localhost:3308)/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	handleRequests()
}
