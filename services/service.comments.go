package services

import (
	"muxblog/dao"
	"muxblog/schemas"
)

type serviceComment struct {
	comments dao.DaoComment
}

func NewServiceComment(comment dao.DaoComment) *serviceComment {
	return &serviceComment{comments: comment}
}

func (s *serviceComment) GetAll() ([]schemas.Comment, error) {
	res, err := s.comments.GetAll()

	return res, err
}

func (s *serviceComment) GetID(id int) (schemas.Comment, error) {

	res, err := s.comments.GetID(id)

	return res, err
}

func (s *serviceComment) Create(input *schemas.Comment) (schemas.Comment, error) {
	var comment schemas.Comment

	comment.IDPOSTCOMMENT = input.IDPOSTCOMMENT
	comment.USERNAMECOMMENT = input.USERNAMECOMMENT
	comment.COMMENT = input.COMMENT

	res, err := s.comments.Create(&comment)

	return res, err
}

func (s *serviceComment) Update(input *schemas.Comment) (schemas.Comment, error) {
	var comment schemas.Comment

	comment.ID = input.ID
	comment.IDPOSTCOMMENT = input.IDPOSTCOMMENT
	comment.USERNAMECOMMENT = input.USERNAMECOMMENT
	comment.COMMENT = input.COMMENT

	res, err := s.comments.Update(&comment)

	return res, err
}

func (s *serviceComment) Delete(id int) error {

	err := s.comments.Delete(id)

	return err
}
