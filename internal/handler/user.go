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

func (h *handler) initUserGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)

		r.Get("/", h.GetUsers)
		r.Get("/{id}", h.GetUser)
		r.Post("/create", h.CreateUser)
		r.Put("/update/{id}", h.UpdateUser)
		r.Delete("/delete/{id}", h.DeleteUser)
	})
}

func (h *handler) GetUsers(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.User.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get users")
	}

	response.ResponseMessage(w, "Success get users", res, http.StatusOK)

}

func (h *handler) GetUser(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	res, err := h.services.User.FindByID(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get user")
	}

	response.ResponseMessage(w, "Success get user", res, http.StatusOK)

}

func (h *handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := user.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.User.Create(&user)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create user")
		return
	}

	response.ResponseMessage(w, "Success create user", res, http.StatusCreated)
}

func (h *handler) UpdateUser(w http.ResponseWriter, r *http.Request) {

	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	var user request.UserRequest

	user.ID = int(Id)

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := user.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.User.Update(&user)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update user")
		return
	}

	response.ResponseMessage(w, "Success update user", res, http.StatusOK)
}

func (h *handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	err = h.services.User.DeleteId(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete user")
		return
	}

	response.ResponseMessage(w, "Success delete user", nil, http.StatusOK)

}
