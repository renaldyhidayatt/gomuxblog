package dao

import "muxblog/schemas"

type DaoUsers interface {
	GetAll() ([]schemas.Users, error)
	GetBYID(id int) (schemas.Users, error)
	Create(input *schemas.Users) (schemas.Users, error)
	Update(input *schemas.Users) (schemas.Users, error)
	Delete(id int) error
	Login(input *schemas.AuthLogin) (schemas.Users, error)
}
