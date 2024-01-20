package handler

import (
	"encoding/json"
	"muxblog/internal/domain/request"
	"muxblog/internal/domain/response"
	"muxblog/internal/middlewares"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func (h *handler) initCommentGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {

		r.Use(middlewares.MiddlewareAuth)

		r.Get("/", h.GetComments)
		r.Get("/{id}", h.GetComment)
		r.Post("/create", h.CreateComment)
		r.Put("/update/{id}", h.UpdateComment)
		r.Delete("/delete/{id}", h.DeleteComment)
	})
}

func (h *handler) GetComments(w http.ResponseWriter, r *http.Request) {

	res, err := h.services.Comment.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get comments")
	}

	response.ResponseMessage(w, "Success get comments", res, http.StatusOK)
}

func (h *handler) GetComment(w http.ResponseWriter, r *http.Request) {

	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	res, err := h.services.Comment.FindById(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get comment")
	}

	response.ResponseMessage(w, "Success get comment", res, http.StatusOK)
}

func (h *handler) CreateComment(w http.ResponseWriter, r *http.Request) {
	var comment request.CommentRequest

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := comment.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Comment.Create(&comment)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create comment")
		return
	}

	response.ResponseMessage(w, "Success create comment", res, http.StatusCreated)
}

func (h *handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	var comment request.CommentRequest

	comment.ID = int(Id)

	if err := json.NewDecoder(r.Body).Decode(&comment); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := comment.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Comment.Update(&comment)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update comment")
		return
	}

	response.ResponseMessage(w, "Success update comment", res, http.StatusOK)
}

func (h *handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	err = h.services.Comment.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete comment")
		return
	}

	response.ResponseMessage(w, "Success delete comment", nil, http.StatusOK)
}
