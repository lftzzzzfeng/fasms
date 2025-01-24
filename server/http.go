package http

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// ServerParams are the parameters required during server creation.
type ServerParams struct {
	*Config

	Logger  *zap.Logger
	Handler http.Handler
}

// Server contains all the necessary functions to run the http application
// server.
type Server struct {
	logger *zap.Logger
	// these are created internally
	httpServer *http.Server
}

// New creates a new http server to handle incoming http requests.
func New(params *ServerParams) (s *Server, err error) {
	readHeaderTimeout, err := parseDuration(params.Config.ReadHeaderTimeout, "readHeaderTimeout")
	if err != nil {
		return
	}

	idleTimout, err := parseDuration(params.Config.IdleTimeout, "idleTimout")
	if err != nil {
		return
	}

	httpServer := &http.Server{
		Addr:              ":" + strconv.FormatUint(uint64(params.Config.Port), 10),
		Handler:           params.Handler,
		IdleTimeout:       idleTimout,
		ReadHeaderTimeout: readHeaderTimeout,
	}

	s = &Server{
		logger:     params.Logger,
		httpServer: httpServer,
	}

	return
}

// ListenAndServe always returns a non-nil error. After Shutdown or Close,
// the returned error is http.ErrServerClosed.
func (s *Server) ListenAndServe() error {
	s.logger.Info("HTTP Server started", zap.String("address", s.httpServer.Addr))

	err := s.httpServer.ListenAndServe()
	if !errors.Is(err, http.ErrServerClosed) {
		s.logger.Error("HTTP listener startup failed",
			zap.String("address", s.httpServer.Addr))

		return errors.Wrap(err, "http listener startup failed")
	}

	return http.ErrServerClosed
}

// Shutdown gracefully shuts down the server without interrupting any active
// connections. See net/http Server Shutdown function for details.
func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func parseDuration(dStr, name string) (d time.Duration, err error) {
	d, err = time.ParseDuration(dStr)
	if err != nil {
		err = errors.Wrapf(err, "http server: invalid %s duration - %s", name, dStr)
	}
	return
}
