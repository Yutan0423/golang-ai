
-- name: ListChoisesByQuestionIDs :many
SELECT
    choice_id,
    question_id,
    content
FROM choices
    WHERE q.question_id IN (sqlc.slice('question_ids'));

-- name: InsertChoises :exec
INSERT INTO choices (
    question_id,
    content
) VALUES (?, ?), (?, ?), (?, ?);