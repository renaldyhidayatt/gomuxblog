package handler

import (
	"encoding/json"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/helpers"
	"muxblog/services"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerComment struct {
	service services.CommentService
}

func NewCommentHandler(comment services.CommentService) *handlerComment {
	return &handlerComment{service: comment}
}

func (h *handlerComment) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.FindAll()

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

	res, err := h.service.FindByID(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (db.Comment{}) == res {
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
	var comments request.CommentRequest

	if err := json.NewDecoder(r.Body).Decode(&comments); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Create(&comments)

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

	var comments request.CommentRequest

	if err := json.NewDecoder(r.Body).Decode(&comments); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	comments.ID = int(id)

	res, err := h.service.Update(&comments)

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

	err = h.service.Delete(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusNoContent, "Dota")
}
