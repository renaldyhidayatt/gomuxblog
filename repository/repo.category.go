package repository

import (
	"context"
	"database/sql"
	"log"
	"muxblog/schemas"
)

type categoryRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewCategoryRepository(ctx context.Context, db *sql.DB) *categoryRepository {
	return &categoryRepository{ctx: ctx, db: db}
}

func (r *categoryRepository) GetAll() ([]schemas.Categories, error) {
	var categoriesSchema schemas.Categories
	var categoriesModel []schemas.Categories

	row, err := r.db.QueryContext(r.ctx, "SELECT * FROM categories ORDER BY id ASC")

	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		if err = row.Scan(&categoriesSchema.ID, &categoriesSchema.Name); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		categoriesModel = append(categoriesModel, categoriesSchema)
	}

	return categoriesModel, nil
}

func (r *categoryRepository) GetID(id int) (schemas.Categories, error) {
	var categoriesModel schemas.Categories

	row, err := r.db.QueryContext(r.ctx, "SELECT id, name FROM categories WHERE id = ?", id)

	if err != nil {
		log.Fatal("Error Query Category: " + err.Error())
		return categoriesModel, err
	}

	for row.Next() {
		err := row.Scan(&categoriesModel.ID, &categoriesModel.Name)
		if err != nil {
			return categoriesModel, err
		}
	}

	return categoriesModel, nil
}

func (r *categoryRepository) Create(input *schemas.Categories) (schemas.Categories, error) {
	var categoryModel schemas.Categories

	row, err := r.db.PrepareContext(r.ctx, "INSERT INTO categories (name) VALUES (?)")

	if err != nil {
		log.Fatal("Error dibagian create: ", err.Error())
		return categoryModel, err
	}

	res, err := row.Exec(input.Name)

	if err != nil {
		log.Fatal("Error dibagian create execContext : ", err.Error())
		return categoryModel, err
	}

	rowID, err := res.LastInsertId()

	if err != nil {
		log.Fatal("Error dibagian last insertid: ", err.Error())
		return categoryModel, err
	}

	result, err := r.GetID(int(rowID))

	if err != nil {
		log.Fatal(err.Error())
		return categoryModel, err
	}

	return result, nil
}

func (r *categoryRepository) Update(category *schemas.Categories) (schemas.Categories, error) {
	row, err := r.db.PrepareContext(r.ctx, "UPDATE categories SET name=? WHERE id=?")

	var categoryModel schemas.Categories

	if err != nil {
		log.Fatal("Error dibagian update: ", err.Error())
		return categoryModel, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
		return categoryModel, err
	}

	_, queryError := row.ExecContext(r.ctx, category.Name, category.ID)

	if queryError != nil {
		log.Fatal(queryError.Error())
		return categoryModel, err
	}

	res, err := r.GetID(category.ID)

	if err != nil {
		return categoryModel, err
	}

	return res, nil
}

func (r *categoryRepository) Delete(id int) error {

	row, err := r.db.PrepareContext(r.ctx, "DELETE FROM categories WHERE id = ?")

	if err != nil {
		return err
	}

	_, queryError := row.ExecContext(r.ctx, id)

	if queryError != nil {
		return err
	}

	return nil

}
