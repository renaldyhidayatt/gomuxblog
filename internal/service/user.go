package service

import (
	"muxblog/internal/domain/request"
	"muxblog/internal/repository"
	db "muxblog/pkg/database/mysql/sqlc"
	"muxblog/pkg/hash"
	"muxblog/pkg/logger"

	"go.uber.org/zap"
)

type userService struct {
	hash       hash.Hashing
	logger     logger.Logger
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository, hash hash.Hashing, logger logger.Logger) *userService {
	return &userService{repository: repository, hash: hash, logger: logger}
}

func (s *userService) FindAll() (*[]db.GetUsersRow, error) {
	res, err := s.repository.FindAll()

	if err != nil {
		s.logger.Error("failed get users: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *userService) FindByID(id int) (*db.User, error) {
	res, err := s.repository.FindById(id)

	if err != nil {
		s.logger.Error("failed get user by id: ", zap.Error(err))
		return nil, err
	}

	return res, nil
}

func (s *userService) Create(requests *request.UserRequest) (*db.User, error) {
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

func (s *userService) Update(requests *request.UserRequest) (*db.User, error) {
	var userRequest request.UserRequest

	res, err := s.repository.FindById(userRequest.ID)

	if err != nil {
		s.logger.Error("Errror update user: ", zap.Error(err))

		return nil, err
	}

	hashing, err := s.hash.HashPassword(requests.Password)

	if err != nil {
		s.logger.Error("Error hashing password: ", zap.Error(err))

		return nil, err
	}

	userRequest.ID = requests.ID

	userRequest.FirstName = requests.FirstName
	userRequest.LastName = requests.LastName
	userRequest.Email = requests.Email
	userRequest.Password = hashing

	res, err = s.repository.Update(&userRequest)

	if err != nil {
		s.logger.Error("Error update user: ", zap.Error(err))

		return nil, err
	}

	return res, nil
}

func (s *userService) DeleteId(id int) error {
	err := s.repository.Delete(id)

	if err != nil {
		s.logger.Error("Error delete user: ", zap.Error(err))

		return err
	}

	return nil
}
