package main

import (
	"backend/route"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connectDB()
	defer db.Close()
	r := mux.NewRouter()
	r.HandleFunc("/ai", route.AIHandler).Methods("GET")
	r.HandleFunc("/question", route.AIHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() *sql.DB {
	dsn := "user:password@tcp(mysql:3307)/db?charset=utf8&parseTime=true"
	for {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			log.Printf("Error connecting to database: %v", err)
		} else {
			fmt.Println("Successfully connected to database")
			return db
		}
	}
}
