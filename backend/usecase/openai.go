package usecase

import (
	"backend/entity"
	"backend/util"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type OpenAI struct{}

func NewOpenAI() *OpenAI {
	return &OpenAI{}
}

var messages []entity.Message

func (o *OpenAI) ScoreByAnswer(answer string) entity.OpenaiResponse {
	apiKey := util.GetEnv("OPENAI_API_KEY")
	content := fmt.Sprintf(entity.ScoreLastQuestionPrompt, answer)
	messages = append(messages, entity.Message{
		Role:    "user",
		Content: content,
	})
	requestBody := entity.OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	fmt.Printf("requestBody: %v", requestBody)
	requestJson, _ := json.Marshal(requestBody)
	req, err := http.NewRequest("POST", entity.OpenaiURL, bytes.NewBuffer(requestJson))
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	c := &http.Client{
		Timeout: entity.OpenaiTimeout,
	}
	resp, err := c.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

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
	fmt.Printf("res: %v", res)

	messages = append(messages, entity.Message{
		Role:    "assistant",
		Content: res.Choices[0].Messages.Content,
	})
	return res
}
