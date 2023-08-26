package route

import (
	"backend/entity"
	"backend/usecase"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type OpenAI struct {
	usecase usecase.OpenAI
}

func NewOpenAI(uo usecase.OpenAI) *OpenAI {
	return &OpenAI{
		usecase: uo,
	}
}

// func (o *OpenAI) OpenAIHandler(w http.ResponseWriter, r *http.Request) {
// 	res := o.usecase.GetOpenAIResponse()
// 	if len(res.Choices) == 0 {
// 		log.Fatalln("Response not found")
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write([]byte(res.Choices[0].Messages.Content))
// }

func (o *OpenAI) ScoreByAnswer(w http.ResponseWriter, r *http.Request) {
	var input entity.ScoreInput
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	err = json.Unmarshal(body, &input)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	res := o.usecase.ScoreByAnswer(input.Answer)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
