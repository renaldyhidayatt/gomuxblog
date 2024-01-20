package service

import (
	"muxblog/internal/domain/request"
	"muxblog/internal/repository"
	db "muxblog/pkg/database/mysql/sqlc"
	"muxblog/pkg/logger"

	"go.uber.org/zap"
)

type categoryService struct {
	repository repository.CategoryRepository
	logger     logger.Logger
}

func NewCategoryService(repository repository.CategoryRepository, logger logger.Logger) *categoryService {
	return &categoryService{repository: repository, logger: logger}
}

func (s *categoryService) FindAll() (*[]db.Category, error) {
	res, err := s.repository.FindAll()

	if err != nil {
		s.logger.Error("failed get categories: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *categoryService) FindByID(id int) (*db.Category, error) {
	res, err := s.repository.FindByID(id)

	if err != nil {
		s.logger.Error("failed get category by id: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *categoryService) Create(input *request.CategoryRequest) (*db.Category, error) {
	res, err := s.repository.Create(input)

	if err != nil {
		s.logger.Error("failed create category: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *categoryService) Update(input *request.CategoryRequest) (*db.Category, error) {
	res, err := s.repository.Update(input)

	if err != nil {
		s.logger.Error("failed update category: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *categoryService) Delete(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		s.logger.Error("failed delete category: ", zap.Error(err))
		return err
	}

	return nil
}
