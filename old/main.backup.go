/**
* LangChainを使う必要ある時のための予備パッケージ
* 必要な時にすぐに使えるようにするためのメモ
* Ref: https://github.com/tmc/langchaingo/blob/main/docs/docs/getting-started/guide-llm.mdx
**/

package backup

import (
	"context"
	"log"
	"net/http"

	"encoding/json"

	"github.com/gorilla/mux"
	"github.com/tmc/langchaingo/chains"
	"github.com/tmc/langchaingo/llms/openai"
	"github.com/tmc/langchaingo/prompts"
)

type Message struct {
	Content string `json:"content"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/message", MessageHandler).Methods("GET")
	r.HandleFunc("/", HelloHandler).Methods("GET")

	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	model, err := openai.New()
	if err != nil {
		log.Fatal(err)
	}

	template := `
	Webエンジニアの設計力を試す問題を5問出してほしい。
	選択肢は4つにしたい。
	ジャンルは{{.layer}}限定で
	そしてこの5問をデータとして扱えるように以下のようなjson形式で吐き出してもらえますか？
	それ以外の出力はなしでお願い
	[
				{
						"question": "質問1",
						"choises": [
								"選択肢1",
								"選択肢2",
								"選択肢3",
								"選択肢4"
						],
						"answer": "回答1"
				},
				{
						"question": "質問2",
						"choises": [
								"選択肢1",
								"選択肢2",
								"選択肢3",
								"選択肢4"
						],
						"answer": "回答2"
				},
				{
						"question": "質問3",
						"choises": [
								"選択肢1",
								"選択肢2",
								"選択肢3",
								"選択肢4"
						],
						"answer": "回答3"
				}
		]
		`
	prompt := prompts.NewPromptTemplate(
		template,
		[]string{"layer"},
	)
	chain := chains.NewLLMChain(model, prompt)

	res, err := chains.Call(context.Background(), chain, map[string]any{
		"layer": "フロントエンド",
	}, chains.WithModel("gpt-3.5-turbo"))
	if err != nil {
		log.Fatal(err)
	}

	jsonData, err := json.Marshal(res)
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func HelloHandler(w http.ResponseWriter, r *http.Request) {
	message := Message{
		Content: "Hello from gorilla/mux!",
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(message)
}
