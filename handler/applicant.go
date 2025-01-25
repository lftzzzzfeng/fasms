package handler

import (
	"net/http"

	"github.com/lftzzzzfeng/fasms/handler/response"
)

func (r *Router) getAllApplicantHandler(res http.ResponseWriter, req *http.Request) {
	r.Render(http.StatusOK, res, &response.GetAllApplicant{
		Name: "test-name",
	})
}
