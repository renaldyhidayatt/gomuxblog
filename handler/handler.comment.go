package handler

import (
	"encoding/json"
	"muxblog/dao"
	"muxblog/helpers"
	"muxblog/schemas"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerComment struct {
	comment dao.DaoComment
}

func NewCommentHandler(comment dao.DaoComment) *handlerComment {
	return &handlerComment{comment: comment}
}

func (h *handlerComment) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.comment.GetAll()

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerComment) GetID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.comment.GetID(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (schemas.Comment{}) == res {
		helpers.ResponseWithJSON(w, http.StatusNotFound, res)
	} else {
		helpers.ResponseWithJSON(w, http.StatusOK, res)
	}

}

func (h *handlerComment) Create(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	var comments schemas.Comment

	if err := json.NewDecoder(r.Body).Decode(&comments); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.comment.Create(&comments)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusCreated, res)
}

func (h *handlerComment) Update(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	var comments schemas.Comment

	if err := json.NewDecoder(r.Body).Decode(&comments); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	comments.ID = int(id)

	res, err := h.comment.Update(&comments)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerComment) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	err = h.comment.Delete(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusNoContent, "Dota")
}
