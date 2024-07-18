// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: playlists.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const createPlaylist = `-- name: CreatePlaylist :one
INSERT INTO playlists (playlist_id, user_uuid, name)
VALUES ($1, $2, $3)
RETURNING user_uuid, playlist_uuid, playlist_id, name, created_at, updated_at
`

type CreatePlaylistParams struct {
	PlaylistID string    `json:"playlist_id"`
	UserUuid   uuid.UUID `json:"user_uuid"`
	Name       string    `json:"name"`
}

func (q *Queries) CreatePlaylist(ctx context.Context, arg CreatePlaylistParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, createPlaylist, arg.PlaylistID, arg.UserUuid, arg.Name)
	var i Playlist
	err := row.Scan(
		&i.UserUuid,
		&i.PlaylistUuid,
		&i.PlaylistID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const deletePlaylist = `-- name: DeletePlaylist :exec
DELETE FROM playlists
WHERE user_uuid = $1 AND playlist_uuid = $2
`

type DeletePlaylistParams struct {
	UserUuid     uuid.UUID `json:"user_uuid"`
	PlaylistUuid uuid.UUID `json:"playlist_uuid"`
}

func (q *Queries) DeletePlaylist(ctx context.Context, arg DeletePlaylistParams) error {
	_, err := q.db.Exec(ctx, deletePlaylist, arg.UserUuid, arg.PlaylistUuid)
	return err
}

const getPlaylist = `-- name: GetPlaylist :one
SELECT user_uuid, playlist_uuid, playlist_id, name, created_at, updated_at
FROM playlists
WHERE user_uuid = $1 AND playlist_uuid = $2
`

type GetPlaylistParams struct {
	UserUuid     uuid.UUID `json:"user_uuid"`
	PlaylistUuid uuid.UUID `json:"playlist_uuid"`
}

func (q *Queries) GetPlaylist(ctx context.Context, arg GetPlaylistParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, getPlaylist, arg.UserUuid, arg.PlaylistUuid)
	var i Playlist
	err := row.Scan(
		&i.UserUuid,
		&i.PlaylistUuid,
		&i.PlaylistID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getPlaylistUUIDByEventUUID = `-- name: GetPlaylistUUIDByEventUUID :one
Select playlist_uuid
FROM events_playlists_mapping
WHERE event_uuid = $1
`

func (q *Queries) GetPlaylistUUIDByEventUUID(ctx context.Context, eventUuid uuid.UUID) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getPlaylistUUIDByEventUUID, eventUuid)
	var playlist_uuid uuid.UUID
	err := row.Scan(&playlist_uuid)
	return playlist_uuid, err
}

const getPlaylistUUIDByName = `-- name: GetPlaylistUUIDByName :one
SELECT playlist_uuid
FROM playlists
WHERE user_uuid = $1 AND name = $2
`

type GetPlaylistUUIDByNameParams struct {
	UserUuid uuid.UUID `json:"user_uuid"`
	Name     string    `json:"name"`
}

func (q *Queries) GetPlaylistUUIDByName(ctx context.Context, arg GetPlaylistUUIDByNameParams) (uuid.UUID, error) {
	row := q.db.QueryRow(ctx, getPlaylistUUIDByName, arg.UserUuid, arg.Name)
	var playlist_uuid uuid.UUID
	err := row.Scan(&playlist_uuid)
	return playlist_uuid, err
}

const listPlaylists = `-- name: ListPlaylists :many
SELECT user_uuid, playlist_uuid, playlist_id, name, created_at, updated_at
FROM playlists
WHERE user_uuid = $1
`

func (q *Queries) ListPlaylists(ctx context.Context, userUuid uuid.UUID) ([]Playlist, error) {
	rows, err := q.db.Query(ctx, listPlaylists, userUuid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Playlist
	for rows.Next() {
		var i Playlist
		if err := rows.Scan(
			&i.UserUuid,
			&i.PlaylistUuid,
			&i.PlaylistID,
			&i.Name,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePlaylistName = `-- name: UpdatePlaylistName :one
UPDATE playlists
SET name = $1
WHERE user_uuid = $2 AND playlist_uuid = $3
RETURNING user_uuid, playlist_uuid, playlist_id, name, created_at, updated_at
`

type UpdatePlaylistNameParams struct {
	Name         string    `json:"name"`
	UserUuid     uuid.UUID `json:"user_uuid"`
	PlaylistUuid uuid.UUID `json:"playlist_uuid"`
}

func (q *Queries) UpdatePlaylistName(ctx context.Context, arg UpdatePlaylistNameParams) (Playlist, error) {
	row := q.db.QueryRow(ctx, updatePlaylistName, arg.Name, arg.UserUuid, arg.PlaylistUuid)
	var i Playlist
	err := row.Scan(
		&i.UserUuid,
		&i.PlaylistUuid,
		&i.PlaylistID,
		&i.Name,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}
