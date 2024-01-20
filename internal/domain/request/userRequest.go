package request

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

type UserRequest struct {
	ID              int    `json:"id"`
	FirstName       string `json:"firstname"`
	LastName        string `json:"lastname"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func (u *UserRequest) Validate() error {

	if u.Password != u.ConfirmPassword {
		return errors.New("password and confirm_password do not match")
	}

	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}
