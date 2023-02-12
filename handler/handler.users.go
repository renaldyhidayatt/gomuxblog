package handler

import (
	"encoding/json"
	db "muxblog/db/sqlc"
	"muxblog/dto/request"
	"muxblog/helpers"
	"muxblog/services"
	"muxblog/utils"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type handlerUser struct {
	service services.UserServices
}

func NewUserHandler(user services.UserServices) *handlerUser {
	return &handlerUser{service: user}
}

func (h *handlerUser) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.FindAll()

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}
	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerUser) GetBYID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.FindById(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	if (db.User{}) == res {
		helpers.ResponseWithJSON(w, http.StatusNotFound, res)
	} else {
		helpers.ResponseWithJSON(w, http.StatusOK, res)
	}
}

func (h *handlerUser) Create(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	var user request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	res, err := h.service.Create(&user)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusCreated, res)
}

func (h *handlerUser) Update(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	user.ID = int(id)

	res, err := h.service.Update(&user)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)

}

func (h *handlerUser) Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err = h.service.Delete(int(id))

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusNoContent, "Delete")
}

func (h *handlerUser) Login(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}

	var user request.AuthLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	res, err := h.service.Login(&user)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := utils.GenerateJwt(res)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusOK, token)
}
