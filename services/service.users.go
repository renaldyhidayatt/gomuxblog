package services

import (
	"fmt"
	"muxblog/dao"
	"muxblog/schemas"
	"muxblog/utils"
)

type serviceUser struct {
	users dao.DaoUsers
}

func NewServiceUser(user dao.DaoUsers) *serviceUser {
	return &serviceUser{users: user}

}

func (s *serviceUser) GetAll() ([]schemas.Users, error) {
	res, err := s.users.GetAll()

	return res, err
}

func (s *serviceUser) GetBYID(id int) (schemas.Users, error) {

	res, err := s.users.GetBYID(id)

	return res, err
}

func (s *serviceUser) Create(input *schemas.Users) (schemas.Users, error) {
	var users schemas.Users

	users.FirstName = input.FirstName
	users.LastName = input.LastName
	users.Email = input.Email
	users.Password = utils.HashPassword(input.Password)

	res, err := s.users.Create(&users)

	return res, err
}

func (s *serviceUser) Update(input *schemas.Users) (schemas.Users, error) {
	var users schemas.Users

	users.ID = input.ID

	users.FirstName = input.FirstName
	users.LastName = input.LastName
	users.Email = input.Email
	users.Password = input.Password

	fmt.Println(&users)
	res, err := s.users.Update(&users)

	return res, err
}

func (s *serviceUser) Delete(id int) error {
	err := s.users.Delete(id)

	return err
}

func (s *serviceUser) Login(input *schemas.AuthLogin) (schemas.Users, error) {
	var user schemas.AuthLogin

	user.Email = input.Email
	user.Password = input.Password

	res, err := s.users.Login(&user)

	return res, err

}
