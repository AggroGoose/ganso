// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"database/sql"
	"time"
)

type Comment struct {
	ID       int64     `json:"id"`
	UserID   string    `json:"user_id"`
	PostID   string    `json:"post_id"`
	Edited   bool      `json:"edited"`
	DateTime time.Time `json:"date_time"`
	Content  string    `json:"content"`
}

type Permission struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type Post struct {
	ID       string         `json:"id"`
	Slug     sql.NullString `json:"slug"`
	AudioUrl sql.NullString `json:"audio_url"`
}

type PostLike struct {
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type PostSafe struct {
	UserID    string    `json:"user_id"`
	PostID    string    `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}

type Reply struct {
	ID        int64     `json:"id"`
	UserID    string    `json:"user_id"`
	CommentID int64     `json:"comment_id"`
	Edited    bool      `json:"edited"`
	DateTime  time.Time `json:"date_time"`
	Content   string    `json:"content"`
}

type User struct {
	ID          string         `json:"id"`
	Verified    bool           `json:"verified"`
	Banned      bool           `json:"banned"`
	Username    sql.NullString `json:"username"`
	Image       sql.NullString `json:"image"`
	Url         sql.NullString `json:"url"`
	UrlVerified sql.NullBool   `json:"url_verified"`
	CreatedAt   time.Time      `json:"created_at"`
}

type UserPermission struct {
	UserID       string    `json:"user_id"`
	PermissionID int64     `json:"permission_id"`
	CreatedAt    time.Time `json:"created_at"`
}
