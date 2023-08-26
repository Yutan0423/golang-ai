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

type ScoreInput struct {
	Answer string `json:"answer"`
}

const OpenaiURL = "https://api.openai.com/v1/chat/completions"

const OpenaiTimeout = 30 * time.Second

const ScoreLastQuestionPrompt = `
あなたはデータベースの問題を採点しています。
以下の問題の答えについて、10満点で点数をつけて下さい
問題:
銀行のシステムでは、口座残高を管理するためのデータベースがあります。複数のユーザーが同時に取引（入金や出金）を行う可能性があります。2人のユーザーが同時に口座からお金を引き出そうとすると、どのような問題が発生する可能性がありますか？
答え:
%s
出力は以下のようにして下さい
score: 7
`

const QuestionPrompt = `
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
