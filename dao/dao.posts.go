package dao

import "muxblog/schemas"

type DaoPosts interface {
	GetAll() ([]schemas.Post, error)
	GetID(id int) (schemas.Post, error)
	GetIDRelationJoin(id int) (schemas.PostRelationJoin, error)
	Create(input *schemas.Post) (schemas.Post, error)
	Update(input *schemas.Post) (schemas.Post, error)
	Delete(id int) error
}
