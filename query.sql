-- name: GetCliNote :one
SELECT * FROM cli_notes
WHERE id = ? LIMIT 1;

-- name: CreateCliNote :exec
INSERT INTO cli_notes (command, note) VALUES (?, ?);

-- name: GetAll :many
SELECT * from cli_notes
ORDER BY created_at DESC;
