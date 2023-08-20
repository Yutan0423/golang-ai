package route

import (
	"backend/usecase"
	"encoding/json"
	"log"
	"net/http"
)

func AIHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := usecase.GetOpenAIResponse()
	if len(res.Choices) == 0 {
		log.Fatalln("Response not found")
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Choices[0].Messages.Content)
}
