-- name: GetUsers :many
SELECT id, firstname, lastname, email FROM users ORDER BY id DESC;

-- name: GetUser :one
SELECT id, firstname,lastname,email,password FROM users WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (firstname, lastname, email, password) VALUES ($1, $2, $3, $4) RETURNING *;

-- name: UpdateUser :one
UPDATE users SET firstname=$2,lastname=$3,email=$4,password=$5 WHERE id=$1 RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;

-- name: FindByEmailUser :one
SELECT id, firstname, lastname, email, password FROM users WHERE email = $1;