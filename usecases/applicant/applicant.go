package applicant

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/lftzzzzfeng/fasms/handler/request"
	applcrepo "github.com/lftzzzzfeng/fasms/repo/applicant"
	familyrepo "github.com/lftzzzzfeng/fasms/repo/family"
)

type Applicant struct {
	ApplicantRepo applcrepo.Applicant
	FamilyRepo    familyrepo.Family
}

func New(repo applcrepo.Applicant) *Applicant {
	return &Applicant{
		ApplicantRepo: repo,
	}
}

func (a *Applicant) CreateApplicant(ctx context.Context, req *request.CreateApplicant) (
	*domain.Applicant, error) {
	// check existing applicant
	applicant, err := a.GetApplicantByIC(ctx, req.Applicant.IC)
	if err != nil {
		return nil, errors.Wrap(err, "applicantusecases: get applicant by ic failed.")
	}

	if applicant.IC != "" {
		return nil, errors.Wrap(err, "applicantusecases: existing applicant.")
	}

	// create familly
	familyID, err := uuid.NewRandom()
	if err != nil {
		return nil, errors.Wrap(err, "applicantusecases: generate uuid failed.")
	}

	family := &domain.Family{
		ID:      familyID,
		Address: familyrepo.DummyAddress,
	}

	err = a.FamilyRepo.Create(ctx, family)
	if err != nil {
		return nil, errors.Wrap(err, "applicantusecases: create family failed.")
	}

	// create applicant
	for i := 0; i < 1+len(req.Household); i++ {
		applicantID, err := uuid.NewRandom()
		if err != nil {
			return nil, errors.Wrap(err, "applicantusecases: generate uuid failed.")
		}

		newApplicant := &domain.Applicant{
			ID:               applicantID,
			Name:             req.Name,
			Sex:              req.Sex,
			IC:               req.IC,
			FamilyID:         familyID,
			EmploymentStatus: req.EmploymentStatus,
		}

		if i > 0 {
			reqApplc := req.Household[i-1]

			newApplicant = &domain.Applicant{
				ID:               applicantID,
				Name:             reqApplc.Name,
				Sex:              reqApplc.Sex,
				IC:               reqApplc.IC,
				FamilyID:         familyID,
				Relationship:     reqApplc.Relation,
				EmploymentStatus: reqApplc.EmploymentStatus,
			}
		}

		if err := a.ApplicantRepo.Create(ctx, newApplicant); err != nil {
			return nil, errors.Wrap(err, "applicantusecases: create applicant failed.")
		}
	}

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
