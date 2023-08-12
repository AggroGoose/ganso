-- name: CreateComment :one
INSERT INTO comments (
    user_id,
    post_id,
    content
) VALUES ($1, $2, $3) RETURNING *;

-- name: CreateReply :one
INSERT INTO replies (
    user_id,
    comment_id,
    content
) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateComment :one
UPDATE comments
SET content = $2, edited = 'true'
WHERE id = $1
RETURNING *;

-- name: UpdateReply :one
UPDATE replies
SET content = $2, edited = 'true'
WHERE id = $1
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;

-- name: DeleteCommentsByPost :exec
DELETE FROM comments WHERE post_id = $1;

-- name: DeleteCommentsByUser :exec
DELETE FROM comments WHERE user_id = $1;

-- name: DeleteReply :exec
DELETE FROM replies WHERE id = $1;

-- name: DeleteRepliesByPost :exec
DELETE FROM replies WHERE comment_id = $1;

-- name: DeleteRepliesByUser :exec
DELETE FROM replies WHERE user_id = $1;

-- name: GetCommentsForPost :many
SELECT u.username, u.image, c.date_time, c.content, c.id
FROM comments AS c
INNER JOIN users AS u
ON c.user_id=u.id
WHERE post_id = $1
ORDER BY c.id
LIMIT $2
OFFSET $3;

-- name: GetRepliesForComment :many
SELECT u.username, u.image, r.date_time, r.content, r.id
FROM replies AS r
INNER JOIN users AS u 
ON r.user_id=u.id
WHERE comment_id = $1
ORDER BY r.id
LIMIT $2
OFFSET $3;