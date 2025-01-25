package main

import (
	"context"
	"errors"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/jmoiron/sqlx"
	"go.uber.org/zap"

	"github.com/lftzzzzfeng/fasms/config"
	"github.com/lftzzzzfeng/fasms/db"
	pgdb "github.com/lftzzzzfeng/fasms/db/pg"
	"github.com/lftzzzzfeng/fasms/handler"
	applcrepo "github.com/lftzzzzfeng/fasms/repo/applicant"
	familyrepo "github.com/lftzzzzfeng/fasms/repo/family"
	schemerepo "github.com/lftzzzzfeng/fasms/repo/scheme"
	httpserver "github.com/lftzzzzfeng/fasms/server"
	applcux "github.com/lftzzzzfeng/fasms/usecases/applicant"
	schemeux "github.com/lftzzzzfeng/fasms/usecases/scheme"
)

const (
	envKey     = "ENV"
	defaultEnv = "dev"
)

// Must is a helper that wraps a function that returns (any, error) and panics
// if the error is not nil.
func must[C any](h C, err error) C {
	if err != nil {
		panic(err)
	}
	return h
}

// application struct hold the necessary information to bootstrap the application.
type application struct {
	config      *config.Config
	logger      *zap.Logger
	httpServer  *httpserver.Server
	defaultPGDB *sqlx.DB
	execer      db.Execer
	wg          sync.WaitGroup
}

func main() {
	logger := must(zap.NewDevelopment())
	defer logger.Sync()

	app := &application{
		logger: logger,
	}

	app.config = must(config.Load(getEnv(envKey, defaultEnv)))

	app.defaultPGDB = must(pgdb.NewPG(app.config.Database))

	app.execer = must(pgdb.NewExec(app.defaultPGDB, app.logger))

	// repo
	applcRepo := applcrepo.New(app.execer)
	familyRepo := familyrepo.New(app.execer)
	schemeRepo := schemerepo.New(app.execer)

	// usecases
	applcUx := applcux.New(applcRepo, familyRepo)
	schemeUx := schemeux.New(schemeRepo)

	routerConf := &handler.RouterConfig{
		Logger:   app.logger,
		ApplcUx:  applcUx,
		SchemeUx: schemeUx,
	}
	router := handler.New(routerConf)

	app.httpServer = must(httpserver.New(&httpserver.ServerParams{
		Config:  app.config.Server,
		Logger:  app.logger,
		Handler: router.Router(),
	}))

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app.wg.Add(1)
	go func(app *application, ctx context.Context, cancel context.CancelFunc) {
		defer app.wg.Done()
		err := app.httpServer.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			app.shutdown(ctx, cancel)
		}
	}(app, ctx, cancel)

	app.wg.Add(1)
	go func(app *application, ctx context.Context, cancel context.CancelFunc) {
		defer app.wg.Done()
		app.signalHandler(ctx, cancel)
	}(app, ctx, cancel)

	app.wg.Wait()
}

// signalHandler starts application shutdown when a signal is received.
func (app *application) signalHandler(ctx context.Context, cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-sigChan:
		app.logger.Info("Shutdown signal received",
			zap.String("signal", s.String()))
		break
	case <-ctx.Done():
		break
	}

	app.shutdown(ctx, cancel)
}

// shutdown gracefully stops the application.
func (app *application) shutdown(ctx context.Context, cancel context.CancelFunc) {
	defer cancel()
	defer app.defaultPGDB.Close()

	app.logger.Info("Graceful shutdown initialized")

	err := app.httpServer.Shutdown(ctx)
	if err != nil {
		app.logger.Error("HTTP Server shutdown error", zap.Error(err))
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
