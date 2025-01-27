package scheme

import (
	"context"
	"slices"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"github.com/pkg/errors"

	"github.com/lftzzzzfeng/fasms/handler/response"
	applcrepo "github.com/lftzzzzfeng/fasms/repo/applicant"
	crirepo "github.com/lftzzzzfeng/fasms/repo/criterion"
	schemerepo "github.com/lftzzzzfeng/fasms/repo/scheme"
)

type Scheme struct {
	ApplicantRepo applcrepo.Applicant
	SchemeRepo    schemerepo.Scheme
	CriteriRepo   crirepo.Criterion
}

func New(applcRepo applcrepo.Applicant, schemeRepo schemerepo.Scheme,
	criRepo crirepo.Criterion) *Scheme {
	return &Scheme{
		ApplicantRepo: applcRepo,
		SchemeRepo:    schemeRepo,
		CriteriRepo:   criRepo,
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
	applicants, err := s.ApplicantRepo.GetByID(ctx, applcID)
	if err != nil {
		return nil, errors.Wrap(err, "schemeusecases: get applicant failed.")
	}

	if len(applicants) == 0 {
		return nil, errors.New("invalid applicant id")
	}

	criterionDetails := []string{}
	criterionIDStrArr := []string{}

	// calculate criteria from household applicants
	for _, applc := range applicants {
		// employment status
		if !slices.Contains(criterionDetails, applc.EmploymentStatus) {
			criterionDetails = append(criterionDetails, applc.EmploymentStatus)
		}

		// citizenship
		firstLetter := string(applc.IC[0])
		if firstLetter == "S" || firstLetter == "T" {
			if !slices.Contains(criterionDetails, crirepo.CRITERION_CITIZEN) {
				criterionDetails = append(criterionDetails, crirepo.CRITERION_CITIZEN)
			}
		}

		// age for primary shcool
		ageLetter := string(applc.IC[1:3])
		age, err := strconv.Atoi(ageLetter)
		if err != nil {
			return nil, errors.Wrap(err, "schemeusecases: get applicant age failed.")
		}
		if age >= 6 && age < 12 {
			if !slices.Contains(criterionDetails, crirepo.CRITERION_PRIMARY_SCHOOL) {
				criterionDetails = append(criterionDetails, crirepo.CRITERION_PRIMARY_SCHOOL)
			}
		}
	}

	// get criteria ids
	criteriaIDs, err := s.CriteriRepo.GetIdsByDetails(ctx, criterionDetails)
	if err != nil {
		return nil, errors.Wrap(err, "schemeusecases: get criteria ids failed.")
	}

	// format criteria ids as string
	for _, id := range criteriaIDs {
		criterionIDStrArr = append(criterionIDStrArr, id.String())
	}

	// get scheme from applicant criteria
	criteriaStr := strings.Join(criterionIDStrArr, "|")
	schemes, err := s.SchemeRepo.GetEligibleSchemesByCritieria(ctx, criteriaStr)
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
