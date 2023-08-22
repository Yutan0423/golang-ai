-- name: InsertAnswer :exec
INSERT INTO answers (
    question_id,
    content
) VALUES (?, ?);