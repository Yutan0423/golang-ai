
-- name: ListQuestionOptionsByQuestionID :many
SELECT
    option_id,
    question_id,
    content
FROM question_options
    WHERE question_id = ?;

-- name: InsertQuestionOptions :exec
INSERT INTO question_options (
    question_id,
    content
) VALUES (?, ?), (?, ?), (?, ?), (?, ?);