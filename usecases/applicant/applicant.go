package applicant

import (
	"context"

	"github.com/pkg/errors"

	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/lftzzzzfeng/fasms/handler/request"
	applcrepo "github.com/lftzzzzfeng/fasms/repo/applicant"
)

type Applicant struct {
	ApplicantRepo applcrepo.Applicant
}

func New(repo applcrepo.Applicant) *Applicant {
	return &Applicant{
		ApplicantRepo: repo,
	}
}

func (a *Applicant) CreateApplicant(ctx context.Context, req *request.CreateApplicant) (
	*domain.Applicant, error) {
	// check existing applicant

	// create familly

	// create applicant

	// id, err := uuid.NewRandom()
	// if err != nil {
	// 	return nil, errors.Wrap(err, "applicantusecases: generate uuid failed.")
	// }

	// applicant := &domain.Applicant{
	// 	ID:               id,
	// 	Name:             name,
	// 	Sex:              sex,
	// 	IC:               ic,
	// 	FamilyID:         familyID,
	// 	Relationship:     relationship,
	// 	EmploymentStatus: employmentStatus,
	// }

	// if err := a.ApplicantRepo.Create(ctx, applicant); err != nil {
	// 	return nil, errors.Wrap(err, "applicantusecases: create applicant failed.")
	// }

	return nil, nil
}

func (a *Applicant) GetAllApplicants(ctx context.Context) ([]*domain.Applicant, error) {
	applicants, err := a.ApplicantRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "applicationusecases: get all applicants failed.")
	}

	return applicants, nil
}

func (a *Applicant) GetApplicantByIC(ctx context.Context, ic string) (*domain.Applicant, error) {
	applicant, err := a.ApplicantRepo.GetByIC(ctx, ic)
	if err != nil {
		return nil, errors.Wrap(err, "applicationusecases: get applicant failed.")
	}

	return applicant, nil
}
