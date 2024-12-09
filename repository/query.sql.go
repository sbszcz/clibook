// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: query.sql

package repository

import (
	"context"
)

const createCliNote = `-- name: CreateCliNote :exec
INSERT INTO cli_notes (command, note) VALUES (?, ?)
`

type CreateCliNoteParams struct {
	Command string
	Note    string
}

func (q *Queries) CreateCliNote(ctx context.Context, arg CreateCliNoteParams) error {
	_, err := q.db.ExecContext(ctx, createCliNote, arg.Command, arg.Note)
	return err
}

const getAll = `-- name: GetAll :many
SELECT id, command, note, created_at from cli_notes
ORDER BY created_at DESC
`

func (q *Queries) GetAll(ctx context.Context) ([]CliNote, error) {
	rows, err := q.db.QueryContext(ctx, getAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []CliNote
	for rows.Next() {
		var i CliNote
		if err := rows.Scan(
			&i.ID,
			&i.Command,
			&i.Note,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOne = `-- name: GetOne :one
SELECT id, command, note, created_at FROM cli_notes
WHERE id = ? LIMIT 1
`

func (q *Queries) GetOne(ctx context.Context, id int64) (CliNote, error) {
	row := q.db.QueryRowContext(ctx, getOne, id)
	var i CliNote
	err := row.Scan(
		&i.ID,
		&i.Command,
		&i.Note,
		&i.CreatedAt,
	)
	return i, err
}
