package usecase

import (
	"backend/entity"
	"backend/repository"
	"context"
)

type Question struct {
	repository repository.Question
}

func NewQuestion(rq repository.Question) *Question {
	return &Question{
		repository: rq,
	}
}

func (q *Question) CreateQuestions(ctx context.Context, questions []entity.Question) error {
	for _, question := range questions {
		err := q.repository.InsertQuestion(ctx, question)
		if err != nil {
			return err
		}
	}
	return nil
}

func (q *Question) ListQuestions(ctx context.Context) ([]*entity.Question, error) {
	res, err := q.repository.ListQuestions(ctx)
	if err != nil {
		return nil, err
	}
	return res, nil
}
