package scheme

import (
	"context"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/lftzzzzfeng/fasms/handler/response"
	applcrepo "github.com/lftzzzzfeng/fasms/repo/applicant"
	schemerepo "github.com/lftzzzzfeng/fasms/repo/scheme"
)

type Scheme struct {
	ApplicantRepo applcrepo.Applicant
	SchemeRepo    schemerepo.Scheme
}

func New(applcRepo applcrepo.Applicant, schemeRepo schemerepo.Scheme) *Scheme {
	return &Scheme{
		ApplicantRepo: applcRepo,
		SchemeRepo:    schemeRepo,
	}
}

func (s *Scheme) GetAllSchemes(ctx context.Context) ([]*response.GetAllSchemes, error) {
	schemes, err := s.SchemeRepo.GetAll(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "schemeusecases: get all schemes failed.")
	}

	schemesRes := []*response.GetAllSchemes{}

	schemesMap := map[uuid.UUID]*response.GetAllSchemes{}
	if len(schemes) > 0 {
		for _, scheme := range schemes {
			if _, ok := schemesMap[scheme.SchemeID]; !ok {
				schemesMap[scheme.SchemeID] = &response.GetAllSchemes{}
				schemesRes = append(schemesRes, schemesMap[scheme.SchemeID])
			}
		}
	}

	return schemesRes, nil
}

func (s *Scheme) GetEligibleSchemesByApplicant(ctx context.Context, applcID uuid.UUID) (
	[]*response.GetAllSchemes, error) {
	applicant, err := s.ApplicantRepo.GetByID(ctx, applcID)
	if err != nil {
		return nil, errors.Wrap(err, "schemeusecases: get applicant failed.")
	}

	if applicant == nil {
		return nil, errors.New("invalid applicant id")
	}

	return nil, nil
}
