package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	applcux "github.com/lftzzzzfeng/fasms/usecases/applicant"
	appux "github.com/lftzzzzfeng/fasms/usecases/application"
	schemeux "github.com/lftzzzzfeng/fasms/usecases/scheme"
)

type RouterConfig struct {
	Logger   *zap.Logger
	ApplcUx  *applcux.Applicant
	SchemeUx *schemeux.Scheme
	AppUx    *appux.Application
}

type Router struct {
	logger   *zap.Logger
	applcUx  *applcux.Applicant
	schemeUx *schemeux.Scheme
	appUx    *appux.Application
}

func New(routerConf *RouterConfig) *Router {
	return &Router{
		logger:   routerConf.Logger,
		applcUx:  routerConf.ApplcUx,
		schemeUx: routerConf.SchemeUx,
		appUx:    routerConf.AppUx,
	}
}

func (h *Router) Router() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))

	router.Post("/v1/api/applicants", h.createApplicantHandler)
	router.Get("/v1/api/applicants", h.getAllApplicantHandler)

	router.Get("/v1/api/schemes", h.getAllSchemesHandler)

	router.Post("/v1/api/applications", h.createApplicationHandler)

	return router
}
