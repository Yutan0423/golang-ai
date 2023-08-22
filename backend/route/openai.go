package route

import (
	"backend/usecase"
	"log"
	"net/http"
)

func AIHandler(w http.ResponseWriter, r *http.Request) {
	res := usecase.GetOpenAIResponse()
	if len(res.Choices) == 0 {
		log.Fatalln("Response not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(res.Choices[0].Messages.Content))
}

// 後で実装する
// func InsertQuestionsHandler(w http.ResponseWriter, r *http.Request) {
// 	err := usecase.InsertQuestions()
// 	if err != nil {
// 		log.Fatalln("Error: ", err.Error())
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.WriteHeader(http.StatusCreated)
// }