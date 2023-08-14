-- name: CreatePost :one
INSERT INTO posts (
    id
) VALUES ($1) RETURNING *;

-- name: UpdatePostAudio :one
UPDATE posts
SET audio_url = $2
WHERE id = $1
RETURNING *;

-- name: GetPost :one
SELECT * FROM posts
WHERE id = $1 LIMIT 1;

-- name: ListPosts :many
SELECT * FROM posts
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: LikePost :one
INSERT INTO post_likes (
    user_id,
    post_id
) VALUES($1, $2) RETURNING *;

-- name: RemoveLikePost :exec
DELETE FROM post_likes 
WHERE user_id = $1 AND post_id = $2;

-- name: SavePost :one
INSERT INTO post_saves (
    user_id,
    post_id
) VALUES($1, $2) RETURNING *;

-- name: GetUserSaves :many
SELECT * FROM post_saves
WHERE user_id = $1
ORDER BY created_at
LIMIT $2
OFFSET $3;

-- name: GetPostLike :many
SELECT * FROM post_likes
WHERE user_id = $1 AND post_id = $2;

-- name: GetPostSave :many
SELECT * FROM post_saves
WHERE user_id = $1 AND post_id = $2;

-- name: RemoveSavePost :exec
DELETE FROM post_saves AS p 
WHERE p.user_id = $1 AND p.post_id = $2;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;

-- name: RemoveAllUserLikes :exec
DELETE FROM post_likes WHERE user_id = $1;

-- name: RemoveAllUserSaves :exec
DELETE FROM post_saves WHERE user_id = $1;

-- name: PostLikeCount :one
SELECT COUNT(*) FROM post_likes WHERE post_id = $1;

-- name: PostSaveCount :one
SELECT COUNT(*) FROM post_saves WHERE post_id = $1;