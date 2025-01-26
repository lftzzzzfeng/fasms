package handler

import (
	"net/http"

	"github.com/google/uuid"
	"github.com/lftzzzzfeng/fasms/handler/response"
	"go.uber.org/zap"
)

func (r *Router) getAllSchemesHandler(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	schemes, err := r.schemeUx.GetAllSchemes(ctx)
	if err != nil {
		r.logger.Error("error get all schemes", zap.Error(err))

		r.Render(http.StatusOK, res, &response.Error{
			Msg: "error get all schemes",
		})

		return
	}

	r.Render(http.StatusOK, res, schemes)
}

func (r *Router) getEligibleSchemesByApplicant(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	applcIDStr := req.URL.Query().Get("applicant_id")
	applcID, err := uuid.Parse(applcIDStr)
	if err != nil {
		r.Render(http.StatusOK, res, &response.Error{
			Msg: "invalid applicant_id",
		})

		return
	}

	schemes, err := r.schemeUx.GetEligibleSchemesByApplicant(ctx, applcID)
	if err != nil {
		r.Render(http.StatusOK, res, &response.Error{
			Msg: err.Error(),
		})

		return
	}

	r.Render(http.StatusOK, res, schemes)
}
