package repository

import (
	"context"
	"fmt"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
)

type CategoryRepository = interfaces.ICategoryRepository

type categoryRepository struct {
	db      *db.Queries
	context context.Context
}

func NewCategoryRepository(db *db.Queries, context context.Context) *categoryRepository {
	return &categoryRepository{db: db, context: context}
}

func (r *categoryRepository) FindAll() ([]db.Category, error) {
	res, err := r.db.GetCategories(r.context)

	if err != nil {
		return nil, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *categoryRepository) FindByID(id int) (db.Category, error) {
	res, err := r.db.GetCategory(r.context, int32(id))

	if err != nil {
		return db.Category{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil

}

func (r *categoryRepository) Create(input *request.CategoryRequest) (db.Category, error) {

	res, err := r.db.CreateCategory(r.context, input.Name)

	if err != nil {
		return db.Category{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *categoryRepository) Update(input *request.CategoryRequest) (db.Category, error) {
	var categoryRequest db.UpdateCategoryParams
	res, err := r.db.GetCategory(r.context, int32(input.ID))

	if err != nil {
		return db.Category{}, fmt.Errorf("failed error: %w", err)
	}

	categoryRequest.ID = res.ID
	categoryRequest.Name = input.Name

	res, err = r.db.UpdateCategory(r.context, categoryRequest)

	if err != nil {
		return db.Category{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}

func (r *categoryRepository) Delete(id int) error {
	_, err := r.db.GetCategory(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error: %w", err)
	}

	err = r.db.DeleteCategory(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error: %w", err)
	}

	return nil
}
