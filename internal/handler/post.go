package handler

import (
	"fmt"
	"io"
	"muxblog/internal/domain/request"
	"muxblog/internal/domain/response"
	"muxblog/internal/middlewares"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) initPostGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)

		r.Get("/", h.GetPosts)
		r.Get("/{id}", h.GetPost)
		r.Get("/id/{id}", h.GetIdRelationJoin)
		r.Post("/create", h.CreatePost)
		r.Put("/update/{id}", h.UpdatePost)
		r.Delete("/delete/{id}", h.DeletePost)
	})
}

func (h *handler) GetPosts(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.Post.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get posts")
	}

	response.ResponseMessage(w, "Success get posts", res, http.StatusOK)
}

func (h *handler) GetPost(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	res, err := h.services.Post.FindById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get post")
	}

	response.ResponseMessage(w, "Success get post", res, http.StatusOK)
}

func (h *handler) GetIdRelationJoin(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	res, err := h.services.Post.FindByIDRelationJoin(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get post")
	}

	response.ResponseMessage(w, "Success get post", res, http.StatusOK)
}

func (h *handler) CreatePost(w http.ResponseWriter, r *http.Request) {
	var post request.PostRequest

	title := r.FormValue("title")
	slug := r.FormValue("slug")
	file, filename, err := r.FormFile("img") // img
	body := r.FormValue("body")
	category_id := r.FormValue("category_id")
	user_id := r.FormValue("user_id")
	username := r.FormValue("user_name")

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./uploads/%s", filename.Filename))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	category, err := strconv.ParseInt(category_id, 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	user, err := strconv.ParseInt(user_id, 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	post.Title = title
	post.Slug = slug
	post.Img = filename.Filename
	post.Body = body
	post.CategoryID = int(category)
	post.UserID = int(user)
	post.UserName = username

	res, err := h.services.Post.Create(&post)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create post")
	}

	response.ResponseMessage(w, "Success create post", res, http.StatusCreated)

}

func (h *handler) UpdatePost(w http.ResponseWriter, r *http.Request) {

	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	var post request.PostRequest

	title := r.FormValue("title")
	slug := r.FormValue("slug")
	file, filename, err := r.FormFile("img") // img
	body := r.FormValue("body")
	category_id := r.FormValue("category_id")
	user_id := r.FormValue("user_id")
	username := r.FormValue("user_name")

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./uploads/%s", filename.Filename))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	defer dst.Close()

	_, err = io.Copy(dst, file)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	category, err := strconv.ParseInt(category_id, 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	user, err := strconv.ParseInt(user_id, 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	post.ID = int(Id)
	post.Title = title
	post.Slug = slug
	post.Img = filename.Filename
	post.Body = body
	post.CategoryID = int(category)
	post.UserID = int(user)
	post.UserName = username

	res, err := h.services.Post.Update(&post)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update post")
	}

	response.ResponseMessage(w, "Success update post", res, http.StatusOK)
}

func (h *handler) DeletePost(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	err = h.services.Post.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete post")
		return
	}

	response.ResponseMessage(w, "Success delete post", nil, http.StatusOK)
}
