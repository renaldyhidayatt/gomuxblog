package repository

import (
	"context"
	"fmt"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
)

type PostRepository = interfaces.IPostService

type postRepository struct {
	db      *db.Queries
	context context.Context
}

func NewPostsRepository(db *db.Queries, context context.Context) *postRepository {
	return &postRepository{db: db, context: context}
}

func (r *postRepository) FindAll() ([]db.Post, error) {
	res, err := r.db.GetPosts(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *postRepository) FindByID(id int) (db.Post, error) {
	res, err := r.db.GetPost(r.context, int32(id))

	if err != nil {
		return db.Post{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *postRepository) FindByIDRelationJoin(id int) (db.GetPostRelationRow, error) {
	res, err := r.db.GetPostRelation(r.context, int32(id))

	if err != nil {
		return db.GetPostRelationRow{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *postRepository) Create(input *request.PostRequest) (db.Post, error) {
	var postRequest db.CreatePostParams

	postRequest.Title = input.Title
	postRequest.Slug = input.Slug
	postRequest.Img = input.Img
	postRequest.Body = input.Body
	postRequest.CategoryID = int32(input.CategoryID)
	postRequest.UserID = int32(input.UserID)
	postRequest.UserName = input.UserName

	res, err := r.db.CreatePost(r.context, postRequest)

	if err != nil {
		return db.Post{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *postRepository) Update(input *request.PostRequest) (db.Post, error) {
	var postRequest db.UpdatePostParams

	_, err := r.db.GetPost(r.context, int32(input.ID))

	if err != nil {
		return db.Post{}, fmt.Errorf("failed error: %w", err)
	}

	postRequest.Title = input.Title
	postRequest.Slug = input.Slug
	postRequest.Img = input.Img
	postRequest.Body = input.Body
	postRequest.CategoryID = int32(input.CategoryID)
	postRequest.UserID = int32(input.UserID)
	postRequest.UserName = input.UserName

	res, err := r.db.UpdatePost(r.context, postRequest)

	if err != nil {
		return db.Post{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *postRepository) Delete(id int) error {
	_, err := r.db.GetPost(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error: %w", err)
	}

	err = r.db.DeletePost(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error: %w", err)
	}

	return nil

}
