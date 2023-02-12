-- name: GetPosts :many
SELECT * FROM posts ORDER BY id DESC;

-- name: GetPostRelation :one
SELECT posts.id AS post_id, posts.title, comments.id AS comment_id, comments.id_post_comment, comments.user_name_comment,comments.comment FROM comments JOIN posts ON posts.id = comments.id_post_comment WHERE posts.id = $1;

-- name: GetPost :one
SELECT id,title,slug,img,body,category_id,user_id,user_name FROM posts WHERE id = $1;


-- name: CreatePost :one
INSERT INTO posts (title,slug,img,body,category_id,user_id,user_name) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING *;

-- name: UpdatePost :one
UPDATE posts SET title=$2,slug=$3,img=$4,body=$5,category_id=$6,user_id=$7,user_name=$8 WHERE id=$1 RETURNING *;

-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1;