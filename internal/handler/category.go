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

func (h *handler) initCategoryGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Use(middlewares.MiddlewareAuth)

		r.Get("/", h.GetCategories)
		r.Get("/{id}", h.GetCategory)
		r.Post("/create", h.CreateCategory)
		r.Put("/update/{id}", h.UpdateCategory)
		r.Delete("/delete/{id}", h.DeleteCategory)
	})
}

func (h *handler) GetCategories(w http.ResponseWriter, r *http.Request) {
	res, err := h.services.Category.FindAll()

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get categories")
	}

	response.ResponseMessage(w, "Success get categories", res, http.StatusOK)
}

func (h *handler) GetCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	res, err := h.services.Category.FindByID(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error get category")
	}

	response.ResponseMessage(w, "Success get category", res, http.StatusOK)
}

func (h *handler) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var category request.CategoryRequest

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := category.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Category.Create(&category)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create category")
		return
	}

	response.ResponseMessage(w, "Success create category", res, http.StatusCreated)
}

func (h *handler) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	var category request.CategoryRequest

	category.ID = int(Id)

	if err := json.NewDecoder(r.Body).Decode(&category); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := category.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Category.Update(&category)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error update category")
		return
	}

	response.ResponseMessage(w, "Success update category", res, http.StatusOK)
}

func (h *handler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	Id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid id")
	}

	err = h.services.Category.Delete(int(Id))

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error delete category")
		return
	}

	response.ResponseMessage(w, "Success delete category", nil, http.StatusOK)
}
