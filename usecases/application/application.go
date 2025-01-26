package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/lftzzzzfeng/fasms/handler/request"
	apprepo "github.com/lftzzzzfeng/fasms/repo/application"
	"github.com/pkg/errors"
)

type Application struct {
	AppRepo apprepo.Application
}

func New(appRepo apprepo.Application) *Application {
	return &Application{
		AppRepo: appRepo,
	}
}

func (a *Application) CreateApplication(ctx context.Context, req *request.CreateApplication) error {
	// check existing application
	application, err := a.AppRepo.GetByApplcIDAndSchemeID(ctx, req.ApplcID, req.SchemeID)
	if err != nil {
		return errors.Wrap(err, "applicationusecases: get app by applc_id and scheme_id failed.")
	}

	if application != nil {
		return errors.Wrap(err, "applicationusecases: existing application.")
	}

	// create application
	appID, err := uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "applicationusecases: generate uuid failed.")
	}

	application = &domain.Application{
		ID:          appID,
		ApplicantID: req.ApplcID,
		SchemeID:    req.SchemeID,
	}

	err = a.AppRepo.Create(ctx, application)
	if err != nil {
		return errors.Wrap(err, "applicationusecases: create application failed.")
	}

	return nil
}
