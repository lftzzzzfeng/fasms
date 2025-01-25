package handler

import "net/http"

func (r *Router) getAllApplicantHandler(res http.ResponseWriter, req *http.Request) {
	r.Render(http.StatusOK, res, `{"key":"value}`)
}
