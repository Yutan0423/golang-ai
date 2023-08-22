// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: answers.sql

package sqlc

import (
	"context"
)

const insertAnswer = `-- name: InsertAnswer :exec
INSERT INTO answers (
    question_id,
    content
) VALUES (?, ?)
`

type InsertAnswerParams struct {
	QuestionID int32
	Content    string
}

func (q *Queries) InsertAnswer(ctx context.Context, arg InsertAnswerParams) error {
	_, err := q.db.ExecContext(ctx, insertAnswer, arg.QuestionID, arg.Content)
	return err
}
