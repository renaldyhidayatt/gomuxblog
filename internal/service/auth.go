package service

import (
	"muxblog/internal/domain/request"
	"muxblog/internal/repository"
	"muxblog/pkg/auth"
	db "muxblog/pkg/database/mysql/sqlc"
	"muxblog/pkg/hash"
	"muxblog/pkg/logger"

	"go.uber.org/zap"
)

type authService struct {
	hash       hash.Hashing
	repository repository.UserRepository
	token      auth.TokenManager
	logger     logger.Logger
}

func NewAuthService(repository repository.UserRepository, hash hash.Hashing, token auth.TokenManager, logger logger.Logger) *authService {
	return &authService{repository: repository, hash: hash, token: token, logger: logger}
}

func (s *authService) Register(requests *request.UserRequest) (*db.User, error) {
	var createRequest request.UserRequest

	hashing, err := s.hash.HashPassword(requests.Password)

	if err != nil {
		s.logger.Error("Error hashing password: ", zap.Error(err))
		return nil, err
	}

	_, err = s.repository.FindByEmail(requests.Email)

	if err != nil {
		s.logger.Error("Error fetching user: ", zap.Error(err))

		return nil, err
	}

	createRequest.FirstName = requests.FirstName
	createRequest.LastName = requests.LastName
	createRequest.Email = requests.Email
	createRequest.Password = hashing

	user, err := s.repository.Create(&createRequest)

	if err != nil {
		s.logger.Error("Error creating user: ", zap.Error(err))
		return nil, err
	}

	return user, nil

}

func (s *authService) Login(request *request.AuthLoginRequest) (string, error) {
	res, err := s.repository.FindByEmail(request.Email)
	if err != nil {
		s.logger.Error("failed login: ", zap.Error(err))
		return "", err
	}

	err = s.hash.ComparePassword(res.Password, request.Password)

	if err != nil {
		s.logger.Error("Error comparing password: ", zap.Error(err))

		return "", err
	}

	token, err := s.createJwt(int(res.ID))

	if err != nil {
		s.logger.Error("failed create jwt token: ", zap.Error(err))
		return "", err
	}

	return token, nil
}

func (s *authService) createJwt(id int) (string, error) {
	token, err := s.token.NewJwtToken(id)
	if err != nil {
		s.logger.Error("failed create jwt token: ", zap.Error(err))
		return "", err
	}
	return token, nil
}
