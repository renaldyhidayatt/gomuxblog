package schemas

type Comment struct {
	ID              int    `json:"id"`
	IDPOSTCOMMENT   int    `json:"id_post_comment"`
	USERNAMECOMMENT string `json:"user_name_comment" `
	COMMENT         string `json:"comment"`
}
