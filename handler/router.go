package handler

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"

	applcux "github.com/lftzzzzfeng/fasms/usecases/applicant"
)

type HandlerConfig struct {
	Logger  *zap.Logger
	ApplcUx *applcux.Applicant
}

type Handler struct {
	logger  *zap.Logger
	applcUx *applcux.Applicant
}

func New(hdlConf *HandlerConfig) *Handler {
	return &Handler{
		logger:  hdlConf.Logger,
		applcUx: hdlConf.ApplcUx,
	}
}

func (h *Handler) Router() (chi.Router, error) {
	router := chi.NewRouter()

	router.Use(middleware.AllowContentType("application/json"))

	router.Get("/api/applicants", h.applicantHandler)

	return router, nil
}
