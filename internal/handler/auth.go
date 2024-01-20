package handler

import (
	"encoding/json"
	"muxblog/internal/domain/request"
	"muxblog/internal/domain/response"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *handler) initAuthGroup(prefix string, r *chi.Mux) {
	r.Route(prefix, func(r chi.Router) {
		r.Get("/", h.handlerHello)
		r.Post("/login", h.handleLogin)
		r.Post("/register", h.handleRegister)
	})
}

func (h *handler) handlerHello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello"))
}

func (h *handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	var register request.UserRequest

	if err := json.NewDecoder(r.Body).Decode(&register); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := register.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Auth.Register(&register)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error create user")
		return
	}

	response.ResponseMessage(w, "Success create user", res, http.StatusCreated)
}

func (h *handler) handleLogin(w http.ResponseWriter, r *http.Request) {
	var login request.AuthLoginRequest

	if err := json.NewDecoder(r.Body).Decode(&login); err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error invalid request")
	}

	if err := login.Validate(); err != nil {
		response.ResponseError(w, http.StatusBadRequest, "Error invalid validate request")
		return
	}

	res, err := h.services.Auth.Login(&login)

	if err != nil {
		response.ResponseError(w, http.StatusInternalServerError, "Error login user")
		return
	}

	response.ResponseToken(w, "Success login", res, http.StatusOK)
}
