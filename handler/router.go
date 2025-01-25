package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	applcux "github.com/lftzzzzfeng/fasms/usecases/applicant"
)

type RouterConfig struct {
	Logger  *zap.Logger
	ApplcUx *applcux.Applicant
}

type Router struct {
	logger  *zap.Logger
	applcUx *applcux.Applicant
}

func New(routerConf *RouterConfig) *Router {
	return &Router{
		logger:  routerConf.Logger,
		applcUx: routerConf.ApplcUx,
	}
}

func (h *Router) Router() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))

	router.Post("/v1/api/applicants", h.createApplicantHandler)
	router.Get("/v1/api/applicants", h.getAllApplicantHandler)

	router.Get("/v1/api/schemes", h.getAllApplicantHandler)

	return router
}
