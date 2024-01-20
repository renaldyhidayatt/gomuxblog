package repository

import (
	"context"
	db "muxblog/pkg/database/mysql/sqlc"
)

type Repositories struct {
	Category CategoryRepository
	Post     PostRepository
	Comment  CommentRepository
	User     UserRepository
}

func NewRepositories(db *db.Queries, context context.Context) *Repositories {
	return &Repositories{
		Category: NewCategoryRepository(db, context),
		Post:     NewPostsRepository(db, context),
		Comment:  NewCommentRepository(db, context),
		User:     NewUserRepository(db, context),
	}
}
