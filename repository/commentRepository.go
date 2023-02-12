package repository

import (
	"context"
	"fmt"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
)

type CommentRepository = interfaces.ICommentRepository

type commentRepository struct {
	db      *db.Queries
	context context.Context
}

func NewCommentRepository(db *db.Queries, context context.Context) *commentRepository {
	return &commentRepository{db: db, context: context}
}

func (r *commentRepository) FindAll() ([]db.Comment, error) {
	res, err := r.db.GetComments(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *commentRepository) FindByID(id int) (db.Comment, error) {
	res, err := r.db.GetComment(r.context, int32(id))

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *commentRepository) Create(input *request.CommentRequest) (db.Comment, error) {
	var commentRequest db.CreateCommentParams

	commentRequest.Comment = input.COMMENT
	commentRequest.IDPostComment = int32(input.IDPOSTCOMMENT)
	commentRequest.UserNameComment = input.USERNAMECOMMENT

	res, err := r.db.CreateComment(r.context, commentRequest)

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *commentRepository) Update(input *request.CommentRequest) (db.Comment, error) {
	var commentRequest db.UpdateCommentParams
	_, err := r.db.GetComment(r.context, int32(input.ID))

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed error: %w", err)
	}
	commentRequest.Comment = input.COMMENT
	commentRequest.IDPostComment = int32(input.IDPOSTCOMMENT)
	commentRequest.UserNameComment = input.USERNAMECOMMENT

	res, err := r.db.UpdateComment(r.context, commentRequest)

	if err != nil {
		return db.Comment{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *commentRepository) Delete(id int) error {
	_, err := r.db.GetComment(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error: %w", err)
	}

	err = r.db.DeleteComment(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error: %w", err)
	}

	return nil
}
