package services

import (
	"fmt"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
	"muxblog/repository"
)

type PostService = interfaces.IPostService

type postService struct {
	repository repository.PostRepository
}

func NewPostService(repository repository.PostRepository) *postService {
	return &postService{repository: repository}
}

func (s *postService) FindAll() ([]db.Post, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *postService) FindByID(id int) (db.Post, error) {

	res, err := s.repository.FindByID(id)

	return res, err
}

func (s *postService) FindByIDRelationJoin(id int) (db.GetPostRelationRow, error) {
	res, err := s.repository.FindByIDRelationJoin(id)

	return res, err
}

func (s *postService) Create(input *request.PostRequest) (db.Post, error) {
	var post request.PostRequest

	post.Title = input.Title
	post.Slug = input.Slug
	post.Img = input.Img
	post.Body = input.Body
	post.CategoryID = input.CategoryID
	post.UserID = input.UserID
	post.UserName = input.UserName

	res, err := s.repository.Create(&post)

	return res, err
}

func (s *postService) Update(input *request.PostRequest) (db.Post, error) {
	var post request.PostRequest

	post.ID = input.ID
	post.Title = input.Title
	post.Slug = input.Slug
	post.Img = input.Img
	post.Body = input.Body
	post.CategoryID = input.CategoryID
	post.UserID = input.UserID
	post.UserName = input.UserName

	fmt.Println(post)

	res, err := s.repository.Update(&post)

	return res, err
}

func (s *postService) Delete(id int) error {

	err := s.repository.Delete(id)

	return err
}
