package handler

import (
	"net/http"

	"github.com/lftzzzzfeng/fasms/handler/request"
	"github.com/lftzzzzfeng/fasms/handler/response"
	"go.uber.org/zap"
)

func (r *Router) createApplicationHandler(res http.ResponseWriter, req *http.Request) {
	ctx := req.Context()

	app := &request.CreateApplication{}
	if err := r.readJSON(req.Body, app); err != nil {
		r.Render(http.StatusBadRequest, res, nil)
		return
	}

	err := r.appUx.CreateApplication(ctx, app)
	if err != nil {
		r.logger.Error("error create application", zap.Error(err))

		r.Render(http.StatusOK, res, &response.Error{
			Msg: "error create application",
		})

		return
	}

	r.Render(http.StatusCreated, res, app)
}
