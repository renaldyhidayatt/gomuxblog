package repository

import (
	"context"
	"database/sql"
	"log"
	"muxblog/schemas"
)

type postRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewPostsRepository(ctx context.Context, db *sql.DB) *postRepository {
	return &postRepository{ctx: ctx, db: db}
}

func (r *postRepository) GetAll() ([]schemas.Post, error) {
	var postSchema schemas.Post
	var postModel []schemas.Post

	row, err := r.db.QueryContext(r.ctx, "SELECT * FROM posts ORDER BY id DESC")

	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		if err = row.Scan(&postSchema.ID, &postSchema.Title, &postSchema.Slug, &postSchema.Body, &postSchema.CategoryID, &postSchema.UserID, &postSchema.UserName); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		postModel = append(postModel, postSchema)
	}

	return postModel, nil
}

func (r *postRepository) GetIDRelationJoin(id int) (schemas.PostRelationJoin, error) {
	var postrelationModel schemas.PostRelationJoin

	row, err := r.db.QueryContext(r.ctx, "SELECT posts.id AS post_id, posts.title, comments.id AS comment_id, comments.id_post_comment, comments.user_name_comment,comments.comment FROM comments JOIN posts ON posts.id = comments.id_post_comment WHERE posts.id = ?", id)

	if err != nil {
		log.Fatal("Error Query Relation Join " + err.Error())
		return postrelationModel, err
	}

	for row.Next() {
		err := row.Scan(&postrelationModel.Post_id, &postrelationModel.Post_title, &postrelationModel.Comment_id, &postrelationModel.CommentIDPostComment, &postrelationModel.CommentUsername, &postrelationModel.CommentComment)

		if err != nil && err != sql.ErrNoRows {
			return postrelationModel, err
		}

	}

	return postrelationModel, nil
}

func (r *postRepository) GetID(id int) (schemas.Post, error) {
	var postsModel schemas.Post

	row, err := r.db.QueryContext(r.ctx, "SELECT id,title,slug,body,category_id,user_id,user_name FROM posts WHERE id = ?", id)

	if err != nil {
		log.Fatal("Error Query Category: " + err.Error())
		return postsModel, err
	}

	for row.Next() {
		err := row.Scan(&postsModel.ID, &postsModel.Title, &postsModel.Slug, &postsModel.Body, &postsModel.CategoryID, &postsModel.UserID, &postsModel.UserName)
		if err != nil {
			return postsModel, err
		}
	}

	return postsModel, nil
}

func (r *postRepository) Create(input *schemas.Post) (schemas.Post, error) {

	row, err := r.db.PrepareContext(r.ctx, "INSERT INTO posts (title,slug,body,category_id,user_id,user_name) VALUES (?,?,?,?,?,?)")

	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := row.ExecContext(r.ctx, &input.Title, &input.Slug, &input.Body, &input.CategoryID, &input.UserID, &input.UserName)

	if err != nil {
		log.Fatal(err.Error())
	}

	rowID, err := res.LastInsertId()

	if err != nil {
		log.Fatal(err.Error())
	}

	result, err := r.GetID(int(rowID))

	if err != nil {
		log.Fatal(err.Error())
	}

	return result, nil
}

func (r *postRepository) Update(input *schemas.Post) (schemas.Post, error) {
	row, err := r.db.PrepareContext(r.ctx, "UPDATE posts SET title=?,slug=?,body=?,category_id=?,user_id=?,user_name=? WHERE id=?")

	var postModel schemas.Post

	if err != nil {
		return postModel, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}

	_, queryError := row.ExecContext(r.ctx, input.Title, input.Slug, input.Body, input.CategoryID, input.UserID, input.UserName, input.ID)

	if queryError != nil {
		return postModel, err
	}

	res, err := r.GetID(input.ID)

	if err != nil {
		return postModel, err
	}

	return res, nil
}

func (r *postRepository) Delete(id int) error {

	row, err := r.db.PrepareContext(r.ctx, "DELETE FROM posts WHERE id = ?")

	if err != nil {
		return err
	}

	_, queryError := row.ExecContext(r.ctx, id)

	if queryError != nil {
		return err
	}

	return nil

}
