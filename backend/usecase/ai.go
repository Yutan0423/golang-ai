package usecase

import (
	"backend/entity"
	"backend/util"
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

var messages []entity.Message

func GetOpenAIResponse() entity.OpenaiResponse {
	apiKey := util.GetEnv("OPENAI_API_KEY")
	messages = append(messages, entity.Message{
		Role:    "user",
		Content: entity.Prompt,
	})
	requestBody := entity.OpenaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: messages,
	}

	requestJson, _ := json.Marshal(requestBody)

	req, err := http.NewRequest("POST", entity.OpenaiURL, bytes.NewBuffer(requestJson))
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

	messages = append(messages, entity.Message{
		Role:    "assistant",
		Content: res.Choices[0].Messages.Content,
	})
	return res
}
