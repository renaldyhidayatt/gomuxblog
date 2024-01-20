package repository

import (
	"context"
	"errors"
	"muxblog/internal/domain/request"
	db "muxblog/pkg/database/mysql/sqlc"
)

type commentRepository struct {
	db      *db.Queries
	context context.Context
}

func NewCommentRepository(db *db.Queries, context context.Context) *commentRepository {
	return &commentRepository{db: db, context: context}
}

func (r *commentRepository) FindAll() (*[]db.Comment, error) {
	res, err := r.db.GetComments(r.context)

	if err != nil {
		return nil, errors.New("failed get comments")
	}

	return &res, nil
}

func (r *commentRepository) FindByID(id int) (*db.Comment, error) {
	res, err := r.db.GetComment(r.context, int32(id))

	if err != nil {
		return nil, errors.New("failed get comment by id")
	}

	return &res, nil
}

func (r *commentRepository) Create(input *request.CommentRequest) (*db.Comment, error) {
	var commentRequest db.CreateCommentParams

	commentRequest.Comment = input.COMMENT
	commentRequest.IDPostComment = int32(input.IDPOSTCOMMENT)
	commentRequest.UserNameComment = input.USERNAMECOMMENT

	res, err := r.db.CreateComment(r.context, commentRequest)

	if err != nil {
		return nil, errors.New("failed create comment")
	}

	return &res, nil
}

func (r *commentRepository) Update(input *request.CommentRequest) (*db.Comment, error) {
	var commentRequest db.UpdateCommentParams
	_, err := r.db.GetComment(r.context, int32(input.ID))

	if err != nil {
		return nil, errors.New("failed get comment by id")
	}

	commentRequest.Comment = input.COMMENT
	commentRequest.IDPostComment = int32(input.IDPOSTCOMMENT)
	commentRequest.UserNameComment = input.USERNAMECOMMENT

	res, err := r.db.UpdateComment(r.context, commentRequest)

	if err != nil {
		return nil, errors.New("failed update comment")
	}

	return &res, nil
}

func (r *commentRepository) Delete(id int) error {
	_, err := r.db.GetComment(r.context, int32(id))

	if err != nil {
		return errors.New("failed get comment by id")
	}

	err = r.db.DeleteComment(r.context, int32(id))

	if err != nil {
		return errors.New("failed delete comment")
	}

	return nil
}
