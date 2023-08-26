package main

import (
	"backend/repository"
	"backend/route"
	"backend/usecase"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := connectDB()
	defer db.Close()

	questionRepository := repository.NewQuestion(db)
	questionUsecase := usecase.NewQuestion(*questionRepository)
	questionRoute := route.NewQuestion(*questionUsecase)

	openaiUsecase := usecase.NewOpenAI()
	openaiRoute := route.NewOpenAI(*openaiUsecase)

	r := mux.NewRouter()
	r.HandleFunc("/openai", openaiRoute.ScoreByAnswer).Methods("POST")
	r.HandleFunc("/question", questionRoute.ListQuestionsHandler).Methods("GET")
	r.HandleFunc("/question", questionRoute.CreateQuestionsHandler).Methods("POST")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func connectDB() *sqlx.DB {
	dsn := "user:password@tcp(mysql:3306)/db?charset=utf8&parseTime=true"
	for {
		db, err := sqlx.Open("mysql", dsn)
		if err != nil {
			log.Printf("Error connecting to database: %v", err)
		} else {
			fmt.Println("Successfully connected to database")
			return db
		}
	}
}
