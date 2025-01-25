package handler

import (
	"net/http"

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
