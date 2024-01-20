package service

import (
	"fmt"
	"muxblog/internal/domain/request"
	"muxblog/internal/repository"
	db "muxblog/pkg/database/mysql/sqlc"
	"muxblog/pkg/logger"

	"go.uber.org/zap"
)

type postService struct {
	logger     logger.Logger
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository, logger logger.Logger) *postService {
	return &postService{repository: repository, logger: logger}
}

func (s *postService) FindAll() (*[]db.Post, error) {

	res, err := s.repository.FindAll()

	if err != nil {
		s.logger.Error("Error fetching posts: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *postService) FindById(id int) (*db.Post, error) {
	res, err := s.repository.FindByID(id)

	if err != nil {
		s.logger.Error("Error fetching post: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *postService) FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error) {
	res, err := s.repository.FindByIDRelationJoin(id)

	if err != nil {
		s.logger.Error("Error fetching post: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *postService) Create(requests *request.PostRequest) (*db.Post, error) {
	var post request.PostRequest

	post.Title = requests.Title
	post.Slug = requests.Slug
	post.Img = requests.Img
	post.Body = requests.Body
	post.CategoryID = requests.CategoryID
	post.UserID = requests.UserID
	post.UserName = requests.UserName

	res, err := s.repository.Create(&post)

	if err != nil {
		s.logger.Error("Error creating post: ", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (s *postService) Update(requests *request.PostRequest) (*db.Post, error) {
	var post request.PostRequest

	post.ID = requests.ID
	post.Title = requests.Title
	post.Slug = requests.Slug
	post.Img = requests.Img
	post.Body = requests.Body
	post.CategoryID = requests.CategoryID
	post.UserID = requests.UserID
	post.UserName = requests.UserName

	fmt.Println(post)

	res, err := s.repository.Update(&post)

	if err != nil {
		s.logger.Error("Error updating post: ", zap.Error(err))
		return nil, err
	}

	return res, err
}

func (s *postService) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		s.logger.Error("Error delete post: ", zap.Error(err))

		return err
	}

	return nil

}
