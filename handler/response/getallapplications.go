package response

import "github.com/google/uuid"

type GetAllApplications struct {
	ID        uuid.UUID `json:"id"`
	Applicant string    `json:"applicant"`
	Scheme    string    `json:"scheme"`
	AppDate   string    `json:"application_date"`
}
