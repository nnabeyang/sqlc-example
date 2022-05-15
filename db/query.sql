-- name: GetEntries :many
SELECT *
FROM entries
ORDER BY created_at DESC
LIMIT ?;
-- name: CreateEntry :execresult
INSERT INTO entries (content)
VALUES (?);