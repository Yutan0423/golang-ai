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
