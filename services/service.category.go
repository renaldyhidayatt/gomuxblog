package services

import (
	"muxblog/dao"
	"muxblog/schemas"
)

type serviceCategory struct {
	category dao.DaoCategories
}

func NewServiceCategory(category dao.DaoCategories) *serviceCategory {
	return &serviceCategory{category: category}
}

func (s *serviceCategory) GetAll() ([]schemas.Categories, error) {
	res, err := s.category.GetAll()

	return res, err
}

func (s *serviceCategory) GetID(id int) (schemas.Categories, error) {

	res, err := s.category.GetID(id)

	return res, err
}

func (s *serviceCategory) Create(input *schemas.Categories) (schemas.Categories, error) {
	var categories schemas.Categories

	categories.Name = input.Name

	res, err := s.category.Create(&categories)

	return res, err
}

func (s *serviceCategory) Update(input *schemas.Categories) (schemas.Categories, error) {
	var category schemas.Categories

	category.ID = input.ID
	category.Name = input.Name

	res, err := s.category.Update(&category)

	return res, err
}

func (s *serviceCategory) Delete(id int) error {

	err := s.category.Delete(id)

	return err
}
