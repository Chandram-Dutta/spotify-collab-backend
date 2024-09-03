// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: users.sql

package database

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users(email, spotify_id, name)
VALUES ($1, $2, $3)
RETURNING user_uuid, id, created_at, version
`

type CreateUserParams struct {
	Email     interface{} `json:"email"`
	SpotifyID string      `json:"spotify_id"`
	Name      string      `json:"name"`
}

type CreateUserRow struct {
	UserUuid  uuid.UUID          `json:"user_uuid"`
	ID        int64              `json:"id"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
	Version   int32              `json:"version"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (CreateUserRow, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.SpotifyID, arg.Name)
	var i CreateUserRow
	err := row.Scan(
		&i.UserUuid,
		&i.ID,
		&i.CreatedAt,
		&i.Version,
	)
	return i, err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users
WHERE user_uuid = $1
`

func (q *Queries) DeleteUser(ctx context.Context, userUuid uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteUser, userUuid)
	return err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, user_uuid, spotify_id, created_at, updated_at, name, email, activated, version
FROM users
WHERE email = $1
`

func (q *Queries) GetUserByEmail(ctx context.Context, email interface{}) (User, error) {
	row := q.db.QueryRow(ctx, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.SpotifyID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.Activated,
		&i.Version,
	)
	return i, err
}

const getUserBySpotifyID = `-- name: GetUserBySpotifyID :one
SELECT user_uuid
FROM users 
WHERE spotify_id = $1
`

func (q *Queries) GetUserBySpotifyID(ctx context.Context, spotifyID string) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getUserBySpotifyID, spotifyID)
	var user_uuid uuid.UUID
	err := row.Scan(&user_uuid)
	return user_uuid, err
}

const getUserByToken = `-- name: GetUserByToken :one
SELECT tokens.user_uuid, spotify_id
FROM tokens
INNER JOIN users ON users.user_uuid = tokens.user_uuid
WHERE access = $1
`

type GetUserByTokenRow struct {
	UserUuid  uuid.UUID `json:"user_uuid"`
	SpotifyID string    `json:"spotify_id"`
}

func (q *Queries) GetUserByToken(ctx context.Context, access []byte) (GetUserByTokenRow, error) {
	row := q.db.QueryRow(ctx, getUserByToken, access)
	var i GetUserByTokenRow
	err := row.Scan(&i.UserUuid, &i.SpotifyID)
	return i, err
}

const getUserByUUID = `-- name: GetUserByUUID :one
SELECT id, user_uuid, spotify_id, created_at, updated_at, name, email, activated, version
FROM users
WHERE user_uuid = $1
`

func (q *Queries) GetUserByUUID(ctx context.Context, userUuid uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUserByUUID, userUuid)
	var i User
	err := row.Scan(
		&i.ID,
		&i.UserUuid,
		&i.SpotifyID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Email,
		&i.Activated,
		&i.Version,
	)
	return i, err
}
