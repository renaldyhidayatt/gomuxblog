package interfaces

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
)

type IPostRepository interface {
	FindAll() ([]db.Post, error)
	FindByID(id int) (db.Post, error)
	FindByIDRelationJoin(id int) (db.GetPostRelationRow, error)
	Create(input *request.PostRequest) (db.Post, error)
	Update(input *request.PostRequest) (db.Post, error)
	Delete(id int) error
}

type IPostService interface {
	FindAll() ([]db.Post, error)
	FindByID(id int) (db.Post, error)
	FindByIDRelationJoin(id int) (db.GetPostRelationRow, error)
	Create(input *request.PostRequest) (db.Post, error)
	Update(input *request.PostRequest) (db.Post, error)
	Delete(id int) error
}
