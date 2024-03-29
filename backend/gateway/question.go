package gateway

import (
	"backend/entity"
	"context"
)

type Question interface {
	InsertQuestion(ctx context.Context, question entity.Question) error
	ListQuestions(ctx context.Context) ([]*entity.Question, error)
}
