package handler

import (
	"net/http"

	"github.com/lftzzzzfeng/fasms/handler/request"
	"github.com/lftzzzzfeng/fasms/handler/response"
	"go.uber.org/zap"
)

func (r *Router) createApplicantHandler(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	applicant := &request.CreateApplicant{}
	if err := r.readJSON(req.Body, applicant); err != nil {
		r.Render(http.StatusBadRequest, res, nil)
		return
	}

	_, err := r.applcUx.CreateApplicant(ctx, applicant)
	if err != nil {
		r.logger.Error("error create applicant", zap.Error(err))

		r.Render(http.StatusOK, res, &response.Error{
			Msg: "error create applicants",
		})

		return
	}

	r.Render(http.StatusOK, res, applicant)
}

func (r *Router) getAllApplicantHandler(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	applicants, err := r.applcUx.GetAllApplicants(ctx)
	if err != nil {
		r.logger.Error("error get all applicants", zap.Error(err))

		r.Render(http.StatusOK, res, &response.Error{
			Msg: "error get all applicants",
		})

		return
	}

	r.Render(http.StatusOK, res, applicants)
}
