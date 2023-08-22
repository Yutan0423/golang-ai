package gateway

import (
	"context"
)

type OpenAI interface {
	InsertQuestions(ctx context.Context, question string, answer string, choices []string) error
}
