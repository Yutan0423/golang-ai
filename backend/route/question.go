package route

import (
	"backend/entity"
	"backend/usecase"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Question struct {
	usecase usecase.Question
}

func NewQuestion(uq usecase.Question) *Question {
	return &Question{
		usecase: uq,
	}
}

func (q *Question) CreateQuestionsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("curl http://localhost:8080/question -X POST")

	var questions []entity.Question
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &questions)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}
	fmt.Println(questions)

	err = q.usecase.CreateQuestions(context.Background(), questions)
	if err != nil {
		log.Fatalln("Error: ", err.Error())
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}
