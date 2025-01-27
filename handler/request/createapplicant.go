package request

type CreateApplicant struct {
	Applicant
	Household []Household `json:"household,omitempty"`
}

type Household struct {
	Applicant
	Relationship
}

type Applicant struct {
	Name             string `json:"name"`
	IC               string `json:"ic"`
	Sex              string `json:"sex"`
	EmploymentStatus string `json:"employment_status"`
}

type Relationship struct {
	Relation string `json:"relation"`
}
