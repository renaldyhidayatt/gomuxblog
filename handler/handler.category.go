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

type handlerCategory struct {
	service services.CategoryService
}

func NewCategoryHandler(category services.CategoryService) *handlerCategory {
	return &handlerCategory{service: category}
}

func (h *handlerCategory) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.service.FindAll()

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerCategory) GetID(w http.ResponseWriter, r *http.Request) {
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

	if (db.Category{}) == res {
		helpers.ResponseWithJSON(w, http.StatusNotFound, res)
	} else {
		helpers.ResponseWithJSON(w, http.StatusOK, res)
	}

}

func (h *handlerCategory) Create(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	var categories request.CategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	res, err := h.service.Create(&categories)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
	}

	helpers.ResponseWithJSON(w, http.StatusCreated, res)
}

func (h *handlerCategory) Update(w http.ResponseWriter, r *http.Request) {
	if r.Header.Get("Content-Type") != "application/json" {
		http.Error(w, "Gunakan content type application/json", http.StatusBadRequest)
		return
	}
	params := mux.Vars(r)
	id, err := strconv.ParseInt(params["id"], 10, 64)

	if err != nil {
		panic(err.Error())
	}

	var categories request.CategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&categories); err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	categories.ID = int(id)

	res, err := h.service.Update(&categories)

	if err != nil {
		helpers.ResponseWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	helpers.ResponseWithJSON(w, http.StatusOK, res)
}

func (h *handlerCategory) Delete(w http.ResponseWriter, r *http.Request) {
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
