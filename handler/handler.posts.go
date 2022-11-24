package handler

import (
	"fmt"
	"io"
	"muxblog/dao"
	"muxblog/helpers"
	"muxblog/schemas"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerPosts struct {
	posts dao.DaoPosts
}

func NewPostsHandler(posts dao.DaoPosts) *handlerPosts {
	return &handlerPosts{posts: posts}
}

func (h *handlerPosts) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.posts.GetAll()

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerPosts) GetID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.posts.GetID(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (schemas.Post{}) == res {
		helpers.ResponseWithJSON(w, http.StatusNotFound, res)
	} else {
		helpers.ResponseWithJSON(w, http.StatusOK, res)
	}
}

func (h *handlerPosts) GetIDRelationJoin(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.posts.GetIDRelationJoin(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (schemas.PostRelationJoin{}) == res {
		helpers.ResponseWithJSON(w, http.StatusNotFound, res)
	} else {
		helpers.ResponseWithJSON(w, http.StatusOK, res)
	}
}

func (h *handlerPosts) Create(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	slug := r.FormValue("slug")
	file, filename, err := r.FormFile("img") // img
	body := r.FormValue("body")
	category_id := r.FormValue("category_id")
	user_id := r.FormValue("user_id")
	username := r.FormValue("user_name")
	var postModel schemas.Post

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./uploads/%s", filename.Filename))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	category, err := strconv.Atoi(category_id)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := strconv.Atoi(user_id)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	postModel.Title = title
	postModel.Slug = slug
	postModel.Img = filename.Filename
	postModel.Body = body
	postModel.CategoryID = category
	postModel.UserID = user
	postModel.UserName = username

	res, err := h.posts.Create(&postModel)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusCreated, res)
}

func (h *handlerPosts) Update(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	title := r.FormValue("title")
	slug := r.FormValue("slug")
	file, filename, err := r.FormFile("img") // img
	body := r.FormValue("body")
	category_id := r.FormValue("category_id")
	user_id := r.FormValue("user_id")
	username := r.FormValue("user_name")
	var postModel schemas.Post

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	defer file.Close()

	dst, err := os.Create(fmt.Sprintf("./uploads/%s", filename.Filename))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer dst.Close()

	_, err = io.Copy(dst, file)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	category, err := strconv.Atoi(category_id)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	user, err := strconv.Atoi(user_id)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	postModel.ID = int(id)
	postModel.Title = title
	postModel.Slug = slug
	postModel.Img = filename.Filename
	postModel.Body = body
	postModel.CategoryID = category
	postModel.UserID = user
	postModel.UserName = username

	res, err := h.posts.Update(&postModel)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerPosts) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	err = h.posts.Delete(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusNoContent, "Delete")
}
