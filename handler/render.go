package handler

import (
	"bytes"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

const ContentTypeJSON = "application/json; charset=utf-8"

func (h *Router) Render(statusCode int, w http.ResponseWriter, v interface{}) {
	buf := &bytes.Buffer{}
	enc := json.NewEncoder(buf)
	enc.SetEscapeHTML(true)
	if err := enc.Encode(v); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", ContentTypeJSON)
	w.WriteHeader(statusCode)
	_, err := w.Write(buf.Bytes())
	if err != nil {
		h.logger.Error("Writing HTTP response failed", zap.Error(err), zap.Any("data", v))
	}
}
