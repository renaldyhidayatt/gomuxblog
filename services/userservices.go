package services

import (
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
	"muxblog/repository"
)

type UserServices = interfaces.IUserService

type userService struct {
	repository repository.UserRepository
}

func NewUserService(repository repository.UserRepository) *userService {
	return &userService{repository: repository}
}

func (s *userService) FindAll() ([]db.GetUsersRow, error) {
	res, err := s.repository.FindAll()

	return res, err

}

func (s *userService) FindById(id int) (db.User, error) {
	res, err := s.repository.FindById(id)

	return res, err
}

func (s *userService) Create(input *request.UserRequest) (db.User, error) {
	var userRequest request.UserRequest

	userRequest.FirstName = input.FirstName
	userRequest.LastName = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := s.repository.Create(&userRequest)

	return res, err
}

func (s *userService) Update(input *request.UserRequest) (db.User, error) {
	var userRequest request.UserRequest

	userRequest.ID = input.ID

	userRequest.FirstName = input.FirstName
	userRequest.LastName = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := s.repository.Update(&userRequest)

	return res, err
}

func (s *userService) Delete(id int) error {
	err := s.repository.Delete(id)

	return err
}

func (s *userService) Login(input *request.AuthLoginRequest) (db.User, error) {
	var user request.AuthLoginRequest

	user.Email = input.Email
	user.Password = input.Password

	res, err := s.repository.Login(&user)

	return res, err

}
