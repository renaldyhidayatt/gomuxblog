package schemas

type Post struct {
	ID         int    `json:"id" form:"id"`
	Title      string `json:"title" form:"title"`
	Slug       string `json:"slug" form:"slug"`
	Img        string `json:"img" form:"img"`
	Body       string `json:"body" form:"body"`
	CategoryID int    `json:"category_id" form:"category_id"`
	UserID     int    `json:"user_id" form:"user_id"`
	UserName   string `json:"user_name" form:"user_name"`
}

type PostRelationJoin struct {
	Post_id              int    `json:"post_id"`
	Post_title           string `json:"title"`
	Comment_id           int    `json:"comment_id"`
	CommentIDPostComment int    `json:"id_post_comment"`
	CommentUsername      string `json:"user_name_comment"`
	CommentComment       string `json:"comment"`
}
