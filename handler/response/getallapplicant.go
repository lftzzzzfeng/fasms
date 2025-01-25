package response

import "github.com/google/uuid"

type GetAllApplicants struct {
	Applicant
	Household []Household `json:"household,omitempty"`
}

type Household struct {
	Applicant
	Relationship
}

type Applicant struct {
	ID               uuid.UUID `json:"id"`
	Name             string    `json:"name"`
	Sex              string    `json:"sex"`
	IC               string    `json:"ic"`
	EmploymentStatus string    `json:"employment_status"`
	DOB              string    `json:"dob"`
}

type Relationship struct {
	Relation string `json:"relation"`
}
