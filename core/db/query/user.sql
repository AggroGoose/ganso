-- name: CreateUser :one
INSERT INTO users (
    id
) VALUES ($1) RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: GetUserForUpdate :one
SELECT * FROM users 
WHERE id = $1 LIMIT 1
FOR NO KEY UPDATE;

-- name: ListUsers :many
SELECT * FROM users
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateUserIntake :one
UPDATE users
SET username = $2, image = $3, verified = 'true'
WHERE id = $1
RETURNING *;

-- name: UpdateUserName :one
UPDATE users
SET username = $2
WHERE id = $1
RETURNING *;

-- name: UpdateUserImage :one
UPDATE users
SET image = $2
WHERE id = $1
RETURNING *;

-- name: AddUserPermission :one
INSERT INTO user_permissions (
    user_id,
    permission_id
) VALUES ($1, $2) RETURNING *;

-- name: RemoveUserPermission :exec
DELETE FROM user_permissions 
WHERE user_id = $1 AND permission_id = $2;

-- name: RemoveAllPermissionsUser :exec
DELETE FROM user_permissions WHERE user_id = $1;

-- name: RemoveAllPermissionGroup :exec
DELETE FROM user_permissions WHERE permission_id = $1;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;