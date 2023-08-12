// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: post.sql

package db

import (
	"context"
	"database/sql"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (
    id,
    slug
) VALUES ($1, $2) RETURNING id, slug, audio_url
`

type CreatePostParams struct {
	ID   string `json:"id"`
	Slug string `json:"slug"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, createPost, arg.ID, arg.Slug)
	var i Post
	err := row.Scan(&i.ID, &i.Slug, &i.AudioUrl)
	return i, err
}

const getPost = `-- name: GetPost :one
SELECT id, slug, audio_url FROM posts
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPost(ctx context.Context, id string) (Post, error) {
	row := q.db.QueryRowContext(ctx, getPost, id)
	var i Post
	err := row.Scan(&i.ID, &i.Slug, &i.AudioUrl)
	return i, err
}

const likePost = `-- name: LikePost :one
INSERT INTO post_likes (
    user_id,
    post_id
) VALUES($1, $2) RETURNING user_id, post_id, created_at
`

type LikePostParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) LikePost(ctx context.Context, arg LikePostParams) (PostLike, error) {
	row := q.db.QueryRowContext(ctx, likePost, arg.UserID, arg.PostID)
	var i PostLike
	err := row.Scan(&i.UserID, &i.PostID, &i.CreatedAt)
	return i, err
}

const listPosts = `-- name: ListPosts :many
SELECT id, slug, audio_url FROM posts
ORDER BY id
LIMIT $1
OFFSET $2
`

type ListPostsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListPosts(ctx context.Context, arg ListPostsParams) ([]Post, error) {
	rows, err := q.db.QueryContext(ctx, listPosts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Post{}
	for rows.Next() {
		var i Post
		if err := rows.Scan(&i.ID, &i.Slug, &i.AudioUrl); err != nil {
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

const removeAllPostLikes = `-- name: RemoveAllPostLikes :exec
DELETE FROM post_likes AS p WHERE p.post_id = $1
`

func (q *Queries) RemoveAllPostLikes(ctx context.Context, postID string) error {
	_, err := q.db.ExecContext(ctx, removeAllPostLikes, postID)
	return err
}

const removeAllPostSaves = `-- name: RemoveAllPostSaves :exec
DELETE FROM post_saves AS p WHERE p.post_id = $1
`

func (q *Queries) RemoveAllPostSaves(ctx context.Context, postID string) error {
	_, err := q.db.ExecContext(ctx, removeAllPostSaves, postID)
	return err
}

const removeAllUserLikes = `-- name: RemoveAllUserLikes :exec
DELETE FROM post_likes AS p WHERE p.user_id = $1
`

func (q *Queries) RemoveAllUserLikes(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, removeAllUserLikes, userID)
	return err
}

const removeAllUserSaves = `-- name: RemoveAllUserSaves :exec
DELETE FROM post_saves AS p WHERE p.user_id = $1
`

func (q *Queries) RemoveAllUserSaves(ctx context.Context, userID string) error {
	_, err := q.db.ExecContext(ctx, removeAllUserSaves, userID)
	return err
}

const removeLikePost = `-- name: RemoveLikePost :exec
DELETE FROM post_likes 
WHERE user_id = $1 AND post_id = $2
`

type RemoveLikePostParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) RemoveLikePost(ctx context.Context, arg RemoveLikePostParams) error {
	_, err := q.db.ExecContext(ctx, removeLikePost, arg.UserID, arg.PostID)
	return err
}

const removeUserLike = `-- name: RemoveUserLike :exec
DELETE FROM post_likes AS p 
WHERE p.user_id = $1 AND p.post_id = $2
`

type RemoveUserLikeParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) RemoveUserLike(ctx context.Context, arg RemoveUserLikeParams) error {
	_, err := q.db.ExecContext(ctx, removeUserLike, arg.UserID, arg.PostID)
	return err
}

const removeUserSave = `-- name: RemoveUserSave :exec
DELETE FROM post_saves AS p 
WHERE p.user_id = $1 AND p.post_id = $2
`

type RemoveUserSaveParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) RemoveUserSave(ctx context.Context, arg RemoveUserSaveParams) error {
	_, err := q.db.ExecContext(ctx, removeUserSave, arg.UserID, arg.PostID)
	return err
}

const savePost = `-- name: SavePost :one
INSERT INTO post_saves (
    user_id,
    post_id
) VALUES($1, $2) RETURNING user_id, post_id, created_at
`

type SavePostParams struct {
	UserID string `json:"user_id"`
	PostID string `json:"post_id"`
}

func (q *Queries) SavePost(ctx context.Context, arg SavePostParams) (PostSafe, error) {
	row := q.db.QueryRowContext(ctx, savePost, arg.UserID, arg.PostID)
	var i PostSafe
	err := row.Scan(&i.UserID, &i.PostID, &i.CreatedAt)
	return i, err
}

const updatePostAudio = `-- name: UpdatePostAudio :one
UPDATE posts
SET audio_url = $2
WHERE id = $1
RETURNING id, slug, audio_url
`

type UpdatePostAudioParams struct {
	ID       string         `json:"id"`
	AudioUrl sql.NullString `json:"audio_url"`
}

func (q *Queries) UpdatePostAudio(ctx context.Context, arg UpdatePostAudioParams) (Post, error) {
	row := q.db.QueryRowContext(ctx, updatePostAudio, arg.ID, arg.AudioUrl)
	var i Post
	err := row.Scan(&i.ID, &i.Slug, &i.AudioUrl)
	return i, err
}