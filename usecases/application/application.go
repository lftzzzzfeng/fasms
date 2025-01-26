package application

import (
	"context"

	"github.com/lftzzzzfeng/fasms/handler/request"
	apprepo "github.com/lftzzzzfeng/fasms/repo/application"
)

type Application struct {
	AppRepo apprepo.Application
}

func New(appRepo apprepo.Application) *Application {
	return &Application{
		AppRepo: appRepo,
	}
}

func (a *Application) CreateApplication(ctx context.Context, req *request.Application) error {
	// check existing application

	// create application

	return nil
}
