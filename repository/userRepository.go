package repository

import (
	"context"
	"fmt"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/interfaces"
	"muxblog/utils"
)

type UserRepository = interfaces.IUserRepository

type userRepository struct {
	db      *db.Queries
	context context.Context
}

func NewUserRepository(db *db.Queries, context context.Context) *userRepository {
	return &userRepository{db: db, context: context}
}

func (r *userRepository) FindAll() ([]db.GetUsersRow, error) {

	user, err := r.db.GetUsers(r.context)

	if err != nil {
		return nil, fmt.Errorf("")
	}

	return user, nil
}

func (r *userRepository) FindById(id int) (db.User, error) {
	user, err := r.db.GetUser(r.context, int32(id))
	if err != nil {
		return db.User{}, fmt.Errorf("failed")
	}

	return user, nil
}

func (r *userRepository) Create(input *request.UserRequest) (db.User, error) {
	var userRequest db.CreateUserParams

	userRequest.Firstname = input.FirstName
	userRequest.Lastname = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	user, err := r.db.CreateUser(r.context, userRequest)

	if err != nil {
		return db.User{}, fmt.Errorf("failed error")
	}

	return user, nil
}

func (r *userRepository) Update(input *request.UserRequest) (db.User, error) {
	var userRequest db.UpdateUserParams

	resid, err := r.db.GetUser(r.context, int32(input.ID))

	if err != nil {
		return db.User{}, fmt.Errorf("failed error")
	}
	err = utils.ComparePassword(resid.Password, input.Password)

	if err != nil {
		return db.User{}, fmt.Errorf("failed error")
	}

	userRequest.Firstname = input.FirstName
	userRequest.Lastname = input.LastName
	userRequest.Email = input.Email
	userRequest.Password = input.Password

	res, err := r.db.UpdateUser(r.context, userRequest)

	if err != nil {
		return db.User{}, fmt.Errorf("failed error")
	}

	return res, nil
}

func (r *userRepository) Delete(id int) error {
	resid, err := r.db.GetUser(r.context, int32(id))

	if err != nil {
		return fmt.Errorf("failed error")
	}

	err = r.db.DeleteUser(r.context, resid.ID)

	if err != nil {
		return fmt.Errorf("failed error")
	}

	return nil

}

func (r *userRepository) FindByEmail(email string) (db.User, error) {
	row, err := r.db.FindByEmailUser(r.context, email)

	if err != nil {
		return db.User{}, fmt.Errorf("failed error")
	}

	return row, nil
}

func (r *userRepository) Login(input *request.AuthLoginRequest) (db.User, error) {
	res, err := r.FindByEmail(input.Email)

	if err != nil {
		return db.User{}, fmt.Errorf("failed error: %w", err)
	}

	err = utils.ComparePassword(res.Password, input.Password)

	if err != nil {
		return db.User{}, fmt.Errorf("failed error: %w", err)
	}

	return res, nil
}
