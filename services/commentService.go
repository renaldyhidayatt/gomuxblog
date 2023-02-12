package services

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
	"muxblog/repository"
)

type CommentService = interfaces.ICommentService

type commentService struct {
	repository repository.CommentRepository
}

func NewCommentService(repository repository.CommentRepository) *commentService {
	return &commentService{repository: repository}
}

func (s *commentService) FindAll() ([]db.Comment, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *commentService) FindByID(id int) (db.Comment, error) {
	res, err := s.repository.FindByID(id)

	return res, err
}

func (s *commentService) Create(input *request.CommentRequest) (db.Comment, error) {
	var comment request.CommentRequest

	comment.IDPOSTCOMMENT = input.IDPOSTCOMMENT
	comment.USERNAMECOMMENT = input.USERNAMECOMMENT
	comment.COMMENT = input.COMMENT

	res, err := s.repository.Create(&comment)

	return res, err
}

func (s *commentService) Update(input *request.CommentRequest) (db.Comment, error) {
	var comment request.CommentRequest

	comment.ID = input.ID
	comment.IDPOSTCOMMENT = input.IDPOSTCOMMENT
	comment.USERNAMECOMMENT = input.USERNAMECOMMENT
	comment.COMMENT = input.COMMENT

	res, err := s.repository.Update(&comment)

	return res, err
}

func (s *commentService) Delete(id int) error {

	err := s.repository.Delete(id)

	return err
}
