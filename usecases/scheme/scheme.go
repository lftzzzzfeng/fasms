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

	if len(schemes) > 0 {
		schemesMap := map[uuid.UUID]*response.GetAllSchemes{}
		criteriaMap := map[string]*response.Criterion{}
		benefitMap := map[string]*response.Benefit{}

		for _, scheme := range schemes {
			criMapKey := scheme.SchemeID.String() + scheme.CriID.String()
			benfMapKey := scheme.SchemeID.String() + scheme.BnftID.String()

			if _, ok := schemesMap[scheme.SchemeID]; !ok {
				schemesMap[scheme.SchemeID] = &response.GetAllSchemes{
					ID:          scheme.SchemeID,
					Name:        scheme.Name,
					Description: scheme.Description,
				}
				schemesRes = append(schemesRes, schemesMap[scheme.SchemeID])
			}

			if _, ok := criteriaMap[criMapKey]; !ok {
				criteriaMap[criMapKey] = &response.Criterion{
					ID:     scheme.CriID,
					Name:   scheme.Criterion,
					Detail: scheme.CriDetail,
				}
				schemesMap[scheme.SchemeID].Criteria = append(schemesMap[scheme.SchemeID].Criteria, criteriaMap[criMapKey])
			}

			if _, ok := benefitMap[benfMapKey]; !ok {
				benefitMap[benfMapKey] = &response.Benefit{
					ID:     scheme.BnftID,
					Name:   scheme.Benefit,
					Detail: scheme.BenefitDetail,
				}
				schemesMap[scheme.SchemeID].Benefits = append(schemesMap[scheme.SchemeID].Benefits, benefitMap[benfMapKey])
			}
		}
	}

	return schemesRes, nil
}

func (s *Scheme) GetEligibleSchemesByApplicant(ctx context.Context, applcID uuid.UUID) (
	[]*response.GetEligibleScheme, error) {
	applicant, err := s.ApplicantRepo.GetByID(ctx, applcID)
	if err != nil {
		return nil, errors.Wrap(err, "schemeusecases: get applicant failed.")
	}

	if applicant == nil {
		return nil, errors.New("invalid applicant id")
	}

	// get all criteria
	// if applicant.EmploymentStatus == "" {

	// }

	// get scheme from applicant criteria
	schemes, err := s.SchemeRepo.GetEligibleSchemesByCritieria(ctx, "ccfde3cb-9fb4-40cf-8bb5-41b20eded403|8bf66ffe-ee9d-49a2-b8d7-352443bcb755")
	if err != nil {
		return nil, errors.Wrap(err, "schemeusecases: get eligible schemes failed.")
	}

	res := []*response.GetEligibleScheme{}
	for _, entry := range schemes {
		scheme := response.GetEligibleScheme{
			ID:          entry.ID,
			Name:        entry.Name,
			Description: entry.Description,
		}

		res = append(res, &scheme)
	}

	return res, nil
}
