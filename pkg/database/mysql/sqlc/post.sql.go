// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: post.sql

package db

import (
	"context"
)

const createPost = `-- name: CreatePost :one
INSERT INTO posts (title,slug,img,body,category_id,user_id,user_name) VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id, title, slug, img, body, category_id, user_id, user_name
`

type CreatePostParams struct {
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Img        string `json:"img"`
	Body       string `json:"body"`
	CategoryID int32  `json:"category_id"`
	UserID     int32  `json:"user_id"`
	UserName   string `json:"user_name"`
}

func (q *Queries) CreatePost(ctx context.Context, arg CreatePostParams) (Post, error) {
	row := q.queryRow(ctx, q.createPostStmt, createPost,
		arg.Title,
		arg.Slug,
		arg.Img,
		arg.Body,
		arg.CategoryID,
		arg.UserID,
		arg.UserName,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Img,
		&i.Body,
		&i.CategoryID,
		&i.UserID,
		&i.UserName,
	)
	return i, err
}

const deletePost = `-- name: DeletePost :exec
DELETE FROM posts WHERE id = $1
`

func (q *Queries) DeletePost(ctx context.Context, id int32) error {
	_, err := q.exec(ctx, q.deletePostStmt, deletePost, id)
	return err
}

const getPost = `-- name: GetPost :one
SELECT id,title,slug,img,body,category_id,user_id,user_name FROM posts WHERE id = $1
`

func (q *Queries) GetPost(ctx context.Context, id int32) (Post, error) {
	row := q.queryRow(ctx, q.getPostStmt, getPost, id)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Img,
		&i.Body,
		&i.CategoryID,
		&i.UserID,
		&i.UserName,
	)
	return i, err
}

const getPostRelation = `-- name: GetPostRelation :one
SELECT posts.id AS post_id, posts.title, comments.id AS comment_id, comments.id_post_comment, comments.user_name_comment,comments.comment FROM comments JOIN posts ON posts.id = comments.id_post_comment WHERE posts.id = $1
`

type GetPostRelationRow struct {
	PostID          int32  `json:"post_id"`
	Title           string `json:"title"`
	CommentID       int32  `json:"comment_id"`
	IDPostComment   int32  `json:"id_post_comment"`
	UserNameComment string `json:"user_name_comment"`
	Comment         string `json:"comment"`
}

func (q *Queries) GetPostRelation(ctx context.Context, id int32) (GetPostRelationRow, error) {
	row := q.queryRow(ctx, q.getPostRelationStmt, getPostRelation, id)
	var i GetPostRelationRow
	err := row.Scan(
		&i.PostID,
		&i.Title,
		&i.CommentID,
		&i.IDPostComment,
		&i.UserNameComment,
		&i.Comment,
	)
	return i, err
}

const getPosts = `-- name: GetPosts :many
SELECT id, title, slug, img, body, category_id, user_id, user_name FROM posts ORDER BY id DESC
`

func (q *Queries) GetPosts(ctx context.Context) ([]Post, error) {
	rows, err := q.query(ctx, q.getPostsStmt, getPosts)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Post
	for rows.Next() {
		var i Post
		if err := rows.Scan(
			&i.ID,
			&i.Title,
			&i.Slug,
			&i.Img,
			&i.Body,
			&i.CategoryID,
			&i.UserID,
			&i.UserName,
		); err != nil {
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

const updatePost = `-- name: UpdatePost :one
UPDATE posts SET title=$2,slug=$3,img=$4,body=$5,category_id=$6,user_id=$7,user_name=$8 WHERE id=$1 RETURNING id, title, slug, img, body, category_id, user_id, user_name
`

type UpdatePostParams struct {
	ID         int32  `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Img        string `json:"img"`
	Body       string `json:"body"`
	CategoryID int32  `json:"category_id"`
	UserID     int32  `json:"user_id"`
	UserName   string `json:"user_name"`
}

func (q *Queries) UpdatePost(ctx context.Context, arg UpdatePostParams) (Post, error) {
	row := q.queryRow(ctx, q.updatePostStmt, updatePost,
		arg.ID,
		arg.Title,
		arg.Slug,
		arg.Img,
		arg.Body,
		arg.CategoryID,
		arg.UserID,
		arg.UserName,
	)
	var i Post
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Slug,
		&i.Img,
		&i.Body,
		&i.CategoryID,
		&i.UserID,
		&i.UserName,
	)
	return i, err
}
