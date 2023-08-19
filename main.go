package main

import (
	"bytes"
	"fmt"
	"golang-ai/entity"
	"golang-ai/util"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
)

const openaiURL = "https://api.openai.com/v1/chat/completions"

var messages []entity.Message

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/ai", AIHandler).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func AIHandler(w http.ResponseWriter, r *http.Request) {
	messages = append(messages, entity.Message{
		Role:    "user",
		Content: entity.Prompt,
	})
	apiKey := util.GetEnv("OPENAI_API_KEY")

	res := getOpenAIResponse(apiKey)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res.Choices[0].Messages.Content)
}

func getOpenAIResponse(apiKey string) entity.OpenaiResponse {
	requestBody := entity.OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	requestJson, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", openaiURL, bytes.NewBuffer(requestJson))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	c := &http.Client{}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var res entity.OpenaiResponse
	err = json.Unmarshal(body, &res)
	if err != nil {
		println("Error: ", err.Error())
		return entity.OpenaiResponse{}
	}

	fmt.Printf("res: %v", res.Choices)
	messages = append(messages, entity.Message{
		Role:    "assistant",
		Content: res.Choices[0].Messages.Content,
	})
	return res
}
