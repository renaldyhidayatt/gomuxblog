package schemas

type Post struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Body       string `json:"body"`
	CategoryID int    `json:"category_id"`
	UserID     int    `json:"user_id"`
	UserName   string `json:"user_name"`
}

type PostRelationJoin struct {
	Post_id              int    `json:"post_id"`
	Post_title           string `json:"title"`
	Comment_id           int    `json:"comment_id"`
	CommentIDPostComment int    `json:"id_post_comment"`
	CommentUsername      string `json:"user_name_comment"`
	CommentComment       string `json:"comment"`
}
