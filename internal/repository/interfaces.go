package repository

import (
	"muxblog/internal/domain/request"
	db "muxblog/pkg/database/mysql/sqlc"
)

type CategoryRepository interface {
	FindAll() (*[]db.Category, error)
	FindByID(id int) (*db.Category, error)
	Create(input *request.CategoryRequest) (*db.Category, error)
	Update(input *request.CategoryRequest) (*db.Category, error)
	Delete(id int) error
}

type CommentRepository interface {
	FindAll() (*[]db.Comment, error)
	FindByID(id int) (*db.Comment, error)
	Create(input *request.CommentRequest) (*db.Comment, error)
	Update(input *request.CommentRequest) (*db.Comment, error)
	Delete(id int) error
}

type PostRepository interface {
	FindAll() (*[]db.Post, error)
	FindByID(id int) (*db.Post, error)
	FindByIDRelationJoin(id int) (*db.GetPostRelationRow, error)
	Create(input *request.PostRequest) (*db.Post, error)
	Update(input *request.PostRequest) (*db.Post, error)
	Delete(id int) error
}

type UserRepository interface {
	FindAll() (*[]db.GetUsersRow, error)
	FindByEmail(email string) (*db.User, error)
	FindById(id int) (*db.User, error)
	Create(input *request.UserRequest) (*db.User, error)
	Update(input *request.UserRequest) (*db.User, error)
	Delete(id int) error
}
