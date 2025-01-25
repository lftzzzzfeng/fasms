package handler

import (
	"net/http"

	"github.com/lftzzzzfeng/fasms/handler/response"
	"go.uber.org/zap"
)

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
