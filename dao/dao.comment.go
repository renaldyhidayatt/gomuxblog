package dao

import "muxblog/schemas"

type DaoComment interface {
	GetAll() ([]schemas.Comment, error)
	GetID(id int) (schemas.Comment, error)
	Create(input *schemas.Comment) (schemas.Comment, error)
	Update(input *schemas.Comment) (schemas.Comment, error)
	Delete(id int) error
}
