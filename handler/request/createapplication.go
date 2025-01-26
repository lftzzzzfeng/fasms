package request

import "github.com/google/uuid"

type CreateApplication struct {
	ApplcID  uuid.UUID `json:"applicant_id,required"`
	SchemeID uuid.UUID `json:"scheme_id,required"`
}
