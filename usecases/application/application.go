package application

import (
	"context"

	"github.com/google/uuid"
	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/lftzzzzfeng/fasms/handler/request"
	"github.com/lftzzzzfeng/fasms/handler/response"
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

func (a *Application) GetAllApplications(ctx context.Context, offset,
	limit int) ([]*response.GetAllApplications, error) {
	applications, err := a.AppRepo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "applicationusecases: get all app failed.")
	}

	applicationRes := []*response.GetAllApplications{}

	for _, app := range applications {
		app := &response.GetAllApplications{
			ID:        app.ID,
			Applicant: app.ApplcName,
			Scheme:    app.SchemeName,
			AppDate:   app.AppDate,
		}
		applicationRes = append(applicationRes, app)
	}

	return applicationRes, nil
}
