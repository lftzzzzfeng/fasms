package domain

import "github.com/google/uuid"

type Application struct {
	ID          uuid.UUID `db:"id"`
	ApplicantID uuid.UUID `db:"applicant_id"`
	SchemeID    uuid.UUID `db:"scheme_id"`
}
