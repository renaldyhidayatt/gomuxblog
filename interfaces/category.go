package interfaces

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
)

type ICategoryRepository interface {
	FindAll() ([]db.Category, error)
	FindByID(id int) (db.Category, error)
	Create(input *request.CategoryRequest) (db.Category, error)
	Update(input *request.CategoryRequest) (db.Category, error)
	Delete(id int) error
}

type ICategoryService interface {
	FindAll() ([]db.Category, error)
	FindByID(id int) (db.Category, error)
	Create(input *request.CategoryRequest) (db.Category, error)
	Update(input *request.CategoryRequest) (db.Category, error)
	Delete(id int) error
}
