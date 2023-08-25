package entity

import "time"

type OpenaiRequest struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

type OpenaiResponse struct {
	ID      string   `json:"id"`
	Object  string   `json:"object"`
	Created int      `json:"created"`
	Choices []Choice `json:"choices"`
	Usages  Usage    `json:"usage"`
}

type Choice struct {
	Index        int     `json:"index"`
	Messages     Message `json:"message"`
	FinishReason string  `json:"finish_reason"`
}

type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type Usage struct {
	PromptTokens     int `json:"prompt_tokens"`
	CompletionTokens int `json:"completion_tokens"`
	TotalTokens      int `json:"total_tokens"`
}

const OpenaiURL = "https://api.openai.com/v1/chat/completions"

const OpenaiTimeout = 30 * time.Second

const Prompt = `
Webエンジニアの設計力を試す問題を3問出してほしい。
- 選択肢は4つにしたい。
- ジャンルはデータベース限定で
- そしてこの3問をデータとして扱えるように以下のようなjson形式で吐き出してもらえますか？
JSON形式のみの出力でお願いします
[
			{
					"question": "質問1",
					"options": [
						"選択肢1",
						"選択肢2",
						"選択肢3",
						"選択肢4",
					],
					"answer": "回答1"
			},
			{
					"question": "質問2",
					"options": [
						"選択肢1",
						"選択肢2",
						"選択肢3",
						"選択肢4",
					],
					"answer": "回答2"
			},
			{
					"question": "質問3",
					"options": [
						"選択肢1",
						"選択肢2",
						"選択肢3",
						"選択肢4",
					],
					"answer": "回答3"
			}
]
`
