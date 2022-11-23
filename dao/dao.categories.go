package dao

import (
	"muxblog/schemas"
)

type DaoCategories interface {
	GetAll() ([]schemas.Categories, error)
	GetID(id int) (schemas.Categories, error)
	Create(input *schemas.Categories) (schemas.Categories, error)
	Update(input *schemas.Categories) (schemas.Categories, error)
	Delete(id int) error
}
