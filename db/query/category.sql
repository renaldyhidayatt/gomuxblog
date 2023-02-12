-- name: GetCategories :many
SELECT * FROM categories ORDER BY id ASC;

-- name: GetCategory :one
SELECT id, name FROM categories WHERE id = $1 LIMIT 1;

-- name: CreateCategory :one
INSERT INTO categories (name) VALUES ($1) RETURNING *;

-- name: UpdateCategory :one
UPDATE categories SET name=$2 WHERE id=$1 RETURNING *;

-- name: DeleteCategory :exec
DELETE FROM categories WHERE id = $1;