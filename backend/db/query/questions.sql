-- name: ListQuestions :many
SELECT
    q.content AS question_content,
    a.content AS answer_content
FROM questions AS q
INNER JOIN answers AS a ON q.question_id = a.question_id
LIMIT ?;

-- name: InsertQuestion :exec
INSERT INTO questions (
    content
) VALUES (?);

-- name: GetLatestQuestionID :one
SELECT LAST_INSERT_ID();
