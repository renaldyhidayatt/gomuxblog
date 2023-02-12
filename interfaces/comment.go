package interfaces

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
)

type ICommentRepository interface {
	FindAll() ([]db.Comment, error)
	FindByID(id int) (db.Comment, error)
	Create(input *request.CommentRequest) (db.Comment, error)
	Update(input *request.CommentRequest) (db.Comment, error)
	Delete(id int) error
}

type ICommentService interface {
	FindAll() ([]db.Comment, error)
	FindByID(id int) (db.Comment, error)
	Create(input *request.CommentRequest) (db.Comment, error)
	Update(input *request.CommentRequest) (db.Comment, error)
	Delete(id int) error
}
