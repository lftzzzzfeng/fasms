package applicant

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/lftzzzzfeng/fasms/handler/request"
	"github.com/lftzzzzfeng/fasms/handler/response"
	applcrepo "github.com/lftzzzzfeng/fasms/repo/applicant"
	familyrepo "github.com/lftzzzzfeng/fasms/repo/family"
)

type Applicant struct {
	ApplicantRepo applcrepo.Applicant
	FamilyRepo    familyrepo.Family
}

func New(applcRepo applcrepo.Applicant, familyRepo familyrepo.Family,
) *Applicant {
	return &Applicant{
		ApplicantRepo: applcRepo,
		FamilyRepo:    familyRepo,
	}
}

func (a *Applicant) CreateApplicant(ctx context.Context, req *request.CreateApplicant) error {
	// check existing applicant
	applicant, err := a.GetApplicantByIC(ctx, req.Applicant.IC)
	if err != nil {
		return errors.Wrap(err, "applicantusecases: get applicant by ic failed.")
	}

	if applicant != nil {
		return errors.Wrap(err, "applicantusecases: existing applicant.")
	}

	// create familly
	familyID, err := uuid.NewRandom()
	if err != nil {
		return errors.Wrap(err, "applicantusecases: generate uuid failed.")
	}

	family := &domain.Family{
		ID:      familyID,
		Address: familyrepo.DummyAddress,
	}

	err = a.FamilyRepo.Create(ctx, family)
	if err != nil {
		return errors.Wrap(err, "applicantusecases: create family failed.")
	}

	// create applicants
	for i := 0; i < 1+len(req.Household); i++ {
		applicantID, err := uuid.NewRandom()
		if err != nil {
			return errors.Wrap(err, "applicantusecases: generate uuid failed.")
		}

		newApplicant := &domain.Applicant{
			ApplicantCommon: &domain.ApplicantCommon{
				ID:               applicantID,
				Name:             req.Name,
				Sex:              req.Sex,
				IC:               req.IC,
				EmploymentStatus: req.EmploymentStatus,
			},
			FamilyID: familyID,
		}

		if i > 0 {
			reqApplc := req.Household[i-1]

			newApplicant = &domain.Applicant{
				ApplicantCommon: &domain.ApplicantCommon{
					ID:               applicantID,
					Name:             reqApplc.Name,
					Sex:              reqApplc.Sex,
					IC:               reqApplc.IC,
					Relationship:     reqApplc.Relation,
					EmploymentStatus: reqApplc.EmploymentStatus,
				},
				FamilyID: familyID,
			}
		}

		if err := a.ApplicantRepo.Create(ctx, newApplicant); err != nil {
			return errors.Wrap(err, "applicantusecases: create applicant failed.")
		}
	}

	return nil
}

func (a *Applicant) GetAllApplicants(ctx context.Context, offset, limit int) (*response.GetAllApplicants, error) {
	applicants, err := a.ApplicantRepo.GetAll(ctx, offset, limit)
	if err != nil {
		return nil, errors.Wrap(err, "applicationusecases: get all applicants failed.")
	}

	applicantsRes := &response.GetAllApplicants{}
	if len(applicants) > 0 {
		for _, applc := range applicants {
			if applc.Relationship == "" {
				applicantsRes.ID = applc.ID
				applicantsRes.Name = applc.Name
				applicantsRes.Sex = applc.Sex
				applicantsRes.IC = applc.IC
				applicantsRes.EmploymentStatus = applc.EmploymentStatus
			} else {
				var household response.Household
				household.ID = applc.ID
				household.Name = applc.Name
				household.Sex = applc.Sex
				household.IC = applc.IC
				household.EmploymentStatus = applc.EmploymentStatus
				household.Relation = applc.Relationship

				applicantsRes.Household = append(applicantsRes.Household, household)
			}
		}
	}

	return applicantsRes, nil
}

func (a *Applicant) GetApplicantByIC(ctx context.Context, ic string) (*domain.Applicant, error) {
	applicant, err := a.ApplicantRepo.GetByIC(ctx, ic)
	if err != nil {
		return nil, errors.Wrap(err, "applicationusecases: get applicant failed.")
	}

	return applicant, nil
}
