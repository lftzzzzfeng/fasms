package domain

import "github.com/google/uuid"

type SchemeBenefitMapping struct {
	SchemeID  uuid.UUID `db:"scheme_id"`
	BenefitID uuid.UUID `db:"benefit_id"`
}
