package request

import "github.com/go-playground/validator/v10"

type CommentRequest struct {
	ID              int    `json:"id"`
	IDPOSTCOMMENT   int    `json:"id_post_comment"`
	USERNAMECOMMENT string `json:"user_name_comment" `
	COMMENT         string `json:"comment"`
}

func (u *CommentRequest) Validate() error {
	validate := validator.New()

	err := validate.Struct(u)

	if err != nil {
		return err
	}

	return nil
}
