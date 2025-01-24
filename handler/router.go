package handler

import (
	"github.com/go-chi/chi/v5"

	applcux "github.com/lftzzzzfeng/fasms/usecases/applicant"
)

type HandlerConfig struct {
	ApplcUx *applcux.Applicant
}

type Handler struct {
	applcUx *applcux.Applicant
}

func New(hdlConf *HandlerConfig) *Handler {
	return &Handler{
		applcUx: hdlConf.ApplcUx,
	}
}

func (h *Handler) Router() (chi.Router, error) {
	router := chi.NewRouter()

	h.applcUx.CreateApplicant()

	return router, nil
}
