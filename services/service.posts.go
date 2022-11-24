package services

import (
	"fmt"
	"muxblog/dao"
	"muxblog/schemas"
)

type servicePosts struct {
	posts dao.DaoPosts
}

func NewServicePosts(post dao.DaoPosts) *servicePosts {
	return &servicePosts{posts: post}

}

func (s *servicePosts) GetAll() ([]schemas.Post, error) {
	res, err := s.posts.GetAll()

	return res, err
}

func (s *servicePosts) GetID(id int) (schemas.Post, error) {

	res, err := s.posts.GetID(id)

	return res, err
}

func (s *servicePosts) GetIDRelationJoin(id int) (schemas.PostRelationJoin, error) {
	res, err := s.posts.GetIDRelationJoin(id)

	return res, err
}

func (s *servicePosts) Create(input *schemas.Post) (schemas.Post, error) {
	var post schemas.Post

	post.Title = input.Title
	post.Slug = input.Slug
	post.Img = input.Img
	post.Body = input.Body
	post.CategoryID = input.CategoryID
	post.UserID = input.UserID
	post.UserName = input.UserName

	fmt.Println(post)

	res, err := s.posts.Create(&post)

	return res, err
}

func (s *servicePosts) Update(input *schemas.Post) (schemas.Post, error) {
	var post schemas.Post

	post.ID = input.ID
	post.Title = input.Title
	post.Slug = input.Slug
	post.Img = input.Img
	post.Body = input.Body
	post.CategoryID = input.CategoryID
	post.UserID = input.UserID
	post.UserName = input.UserName

	fmt.Println(post)

	res, err := s.posts.Update(&post)

	return res, err
}

func (s *servicePosts) Delete(id int) error {

	err := s.posts.Delete(id)

	return err
}
