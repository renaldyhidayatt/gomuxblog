package interfaces

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
)

type IUserRepository interface {
	FindAll() ([]db.GetUsersRow, error)
	FindById(id int) (db.User, error)
	Create(input *request.UserRequest) (db.User, error)
	Update(input *request.UserRequest) (db.User, error)
	Delete(id int) error
	Login(input *request.AuthLoginRequest) (db.User, error)
}

type IUserService interface {
	FindAll() ([]db.GetUsersRow, error)
	FindById(id int) (db.User, error)
	Create(input *request.UserRequest) (db.User, error)
	Update(input *request.UserRequest) (db.User, error)
	Delete(id int) error
	Login(input *request.AuthLoginRequest) (db.User, error)
}
