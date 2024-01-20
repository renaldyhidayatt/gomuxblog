package repository

import (
	"context"
	"errors"
	"fmt"
	"muxblog/internal/domain/request"
	db "muxblog/pkg/database/mysql/sqlc"
)

type postRepository struct {
	db      *db.Queries
	context context.Context
}

func NewPostsRepository(db *db.Queries, context context.Context) *postRepository {
	return &postRepository{db: db, context: context}
}

func (r *postRepository) FindAll() (*[]db.Post, error) {
	res, err := r.db.GetPosts(r.context)

	if err != nil {
		return nil, errors.New("failed get posts")
	}

	return &res, nil
}

func (r *postRepository) FindByID(id int) (*db.Post, error) {
	res, err := r.db.GetPost(r.context, int32(id))

	if err != nil {
		return nil, fmt.Errorf("failed error: %w", err)
	}

	return &res, nil
}

func (r *postRepository) FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error) {
	res, err := r.db.GetPostRelation(r.context, int32(id))

	if err != nil {
		return nil, errors.New("failed get post by id")
	}

	return &res, nil
}

func (r *postRepository) Create(input *request.PostRequest) (*db.Post, error) {
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
		return nil, errors.New("failed create post")
	}

	return &res, nil
}

func (r *postRepository) Update(input *request.PostRequest) (*db.Post, error) {
	var postRequest db.UpdatePostParams

	_, err := r.db.GetPost(r.context, int32(input.ID))

	if err != nil {
		return nil, errors.New("failed get post by id")
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
		return nil, errors.New("failed update post")
	}

	return &res, nil
}

func (r *postRepository) Delete(id int) error {
	_, err := r.db.GetPost(r.context, int32(id))

	if err != nil {
		return errors.New("failed get post by id")
	}

	err = r.db.DeletePost(r.context, int32(id))

	if err != nil {
		return errors.New("failed delete post")
	}

	return nil

}
