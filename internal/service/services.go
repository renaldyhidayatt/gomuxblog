package service

import (
	"muxblog/internal/repository"
	"muxblog/pkg/auth"
	"muxblog/pkg/hash"
	"muxblog/pkg/logger"
)

type Services struct {
	Auth     AuthService
	User     UserService
	Category CategoryService
	Post     PostService
	Comment  CommentService
}

type Deps struct {
	Repository *repository.Repositories
	Logger     *logger.Logger
	Hash       *hash.Hashing
	Token      auth.TokenManager
}

func NewServices(deps Deps) *Services {
	return &Services{
		Auth:     NewAuthService(deps.Repository.User, *deps.Hash, deps.Token, *deps.Logger),
		User:     NewUserService(deps.Repository.User, *deps.Hash, *deps.Logger),
		Category: NewCategoryService(deps.Repository.Category, *deps.Logger),
		Post:     NewPostService(deps.Repository.Post, *deps.Logger),
		Comment:  NewCommentService(deps.Repository.Comment, *deps.Logger),
	}
}
