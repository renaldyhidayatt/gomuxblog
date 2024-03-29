// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import ()

type Category struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	ID              int32  `json:"id"`
	IDPostComment   int32  `json:"id_post_comment"`
	UserNameComment string `json:"user_name_comment"`
	Comment         string `json:"comment"`
}

type Post struct {
	ID         int32  `json:"id"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Img        string `json:"img"`
	Body       string `json:"body"`
	CategoryID int32  `json:"category_id"`
	UserID     int32  `json:"user_id"`
	UserName   string `json:"user_name"`
}

type User struct {
	ID        int32  `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}
