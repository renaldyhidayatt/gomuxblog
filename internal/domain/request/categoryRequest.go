package request

import (
	"github.com/go-playground/validator/v10"
)

type CategoryRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func (r *CategoryRequest) Validate() error {

	validate := validator.New()

	err := validate.Struct(r)

	if err != nil {
		return err
	}

	return nil

}
