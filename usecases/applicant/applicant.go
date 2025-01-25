package applicant

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/lftzzzzfeng/fasms/domain"
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

func (a *Applicant) CreateApplicant(ctx context.Context, name, sex, ic, relationship,
	employmentStatus string, familyID uuid.UUID) (*domain.Applicant, error) {

	id, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "applicantusecases: generate uuid failed.")
	}

	applicant := &domain.Applicant{
		ID:               id,
		Name:             name,
		Sex:              sex,
		IC:               ic,
		FamilyID:         familyID,
		Relationship:     relationship,
		EmploymentStatus: employmentStatus,
	}

	if err := a.ApplicantRepo.Create(ctx, applicant); err != nil {
		return nil, errors.Wrap(err, "applicantusecases: create applicant failed.")
	}

	return applicant, nil
}

func (a *Applicant) GetAllApplicants(ctx context.Context) ([]*domain.Applicant, error) {

	applicants, err := a.ApplicantRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "applicationusecases: get all applicants failed.")
	}

	return applicants, nil
}
