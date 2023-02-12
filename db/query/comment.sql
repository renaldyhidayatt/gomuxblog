-- name: GetComments :many
SELECT * FROM comments ORDER BY id DESC;

-- name: GetComment :one
SELECT id, id_post_comment,user_name_comment,comment FROM comments WHERE id = $1 LIMIT 1;

-- name: CreateComment :one
INSERT INTO comments (id_post_comment,user_name_comment,comment) VALUES ($1,$2,$3) RETURNING *;


-- name: UpdateComment :one
UPDATE comments SET id_post_comment=$2,user_name_comment=$3,comment=$4 WHERE id=$1 RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;