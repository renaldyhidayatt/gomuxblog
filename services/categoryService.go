package services

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
	"muxblog/repository"
)

type CategoryService = interfaces.ICategoryService

type categoryService struct {
	repository repository.CategoryRepository
}

func NewCategoryService(repository repository.CategoryRepository) *categoryService {
	return &categoryService{repository: repository}
}

func (s *categoryService) FindAll() ([]db.Category, error) {
	res, err := s.repository.FindAll()

	return res, err
}

func (s *categoryService) FindByID(id int) (db.Category, error) {
	res, err := s.repository.FindByID(id)

	return res, err
}

func (s *categoryService) Create(input *request.CategoryRequest) (db.Category, error) {
	res, err := s.repository.Create(input)

	return res, err
}

func (s *categoryService) Update(input *request.CategoryRequest) (db.Category, error) {
	res, err := s.repository.Update(input)

	return res, err
}

func (s *categoryService) Delete(id int) error {
	err := s.repository.Delete(id)

	return err
}
