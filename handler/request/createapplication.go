package request

import "github.com/google/uuid"

type Application struct {
	ApplcID  uuid.UUID `json:"applicant_id"`
	Schemeid uuid.UUID `json:"scheme_id"`
}
