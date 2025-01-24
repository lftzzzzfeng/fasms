package applicant

import (
	"context"
	"fmt"

	"github.com/google/uuid"
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

func (a *Applicant) CreateApplicant() {
	id, _ := uuid.NewRandom()
	fmt.Println("id", id)

	uuid, err := uuid.Parse("11e2a580-6d80-4996-a505-80a7c566eb9c")
	fmt.Println("err", err)

	a.ApplicantRepo.Create(context.Background(), &domain.Applicant{
		ID:               id,
		Name:             "test",
		Sex:              "male",
		IC:               "S89809012G",
		FamilyID:         uuid,
		Relationship:     "son",
		EmploymentStatus: "unemployed",
	})
}
