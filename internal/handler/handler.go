package handler

import (
	"muxblog/internal/service"
	"muxblog/pkg/auth"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type handler struct {
	services     *service.Services
	tokenManager auth.TokenManager
}

func NewHandler(services *service.Services, tokenManager auth.TokenManager) *handler {
	return &handler{
		services:     services,
		tokenManager: tokenManager,
	}
}

func (h *handler) Init() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)

	return router
}

func (h *handler) InitApi(r *chi.Mux) {
	r.Get("/", h.handlerHello)

	h.initAuthGroup("/auth", r)
	h.initUserGroup("/user", r)
	h.initCategoryGroup("/category", r)
	h.initCommentGroup("/comment", r)
}
