package repository

import (
	"backend/db/generated/sqlc"
	"backend/entity"
	"context"

	"github.com/jmoiron/sqlx"
)

type Question struct {
	db *sqlx.DB
}

func NewQuestion(db *sqlx.DB) *Question {
	return &Question{db}
}

func (q *Question) InsertQuestion(ctx context.Context, question entity.Question) error {
	tx, err := q.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	queries := sqlc.New(q.db)

	err = queries.InsertQuestion(ctx, question.Title)
	if err != nil {
		return err
	}

	qid, err := queries.GetLatestQuestionID(ctx)
	if err != nil {
		return err
	}

	err = queries.InsertAnswer(ctx, sqlc.InsertAnswerParams{
		QuestionID: int32(qid),
		Content:    question.Answer,
	})
	if err != nil {
		return err
	}

	err = queries.InsertQuestionOptions(ctx, sqlc.InsertQuestionOptionsParams{
		QuestionID:   int32(qid),
		QuestionID_2: int32(qid),
		QuestionID_3: int32(qid),
		QuestionID_4: int32(qid),
		Content:      question.Options[0],
		Content_2:    question.Options[1],
		Content_3:    question.Options[2],
		Content_4:    question.Options[3],
	})
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (q *Question) ListQuestions(ctx context.Context) ([]*entity.Question, error) {
	queries := sqlc.New(q.db)
	questions, err := queries.ListQuestions(ctx, 3)
	if err != nil {
		return nil, err
	}

	res := make([]*entity.Question, 0, len(questions))
	for _, question := range questions {
		options, err := queries.ListQuestionOptionsByQuestionID(ctx, question.QuestionID)
		if err != nil {
			return nil, err
		}
		ops := make([]string, 0, len(options))
		for _, option := range options {
			ops = append(ops, option.Content)
		}

		res = append(res, &entity.Question{
			Title:   question.QuestionContent,
			Answer:  question.AnswerContent,
			Options: ops,
		})
	}
	return res, nil
}
