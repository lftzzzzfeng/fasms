package applicant

import (
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

}
