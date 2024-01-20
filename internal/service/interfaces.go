package service

import (
	"muxblog/internal/domain/request"
	db "muxblog/pkg/database/mysql/sqlc"
)

type CategoryService interface {
	FindAll() (*[]db.Category, error)
	FindByID(id int) (*db.Category, error)
	Create(input *request.CategoryRequest) (*db.Category, error)
	Update(input *request.CategoryRequest) (*db.Category, error)
	Delete(id int) error
}

type PostService interface {
	FindAll() (*[]db.Post, error)
	FindById(id int) (*db.Post, error)
	FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error)
	Create(requests *request.PostRequest) (*db.Post, error)
	Update(requests *request.PostRequest) (*db.Post, error)
	Delete(id int) error
}

type CommentService interface {
	FindAll() (*[]db.Comment, error)
	FindById(id int) (*db.Comment, error)
	Create(input *request.CommentRequest) (*db.Comment, error)
	Update(input *request.CommentRequest) (*db.Comment, error)
	Delete(id int) error
}

type UserService interface {
	FindAll() (*[]db.GetUsersRow, error)
	FindByID(id int) (*db.User, error)
	Create(requests *request.UserRequest) (*db.User, error)
	Update(requests *request.UserRequest) (*db.User, error)
	DeleteId(id int) error
}

type AuthService interface {
	Register(requests *request.UserRequest) (*db.User, error)
	Login(request *request.AuthLoginRequest) (string, error)
}
