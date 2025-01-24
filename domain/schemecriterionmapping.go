package domain

import "github.com/google/uuid"

type SchemeCriterionMapping struct {
	SchemeID    uuid.UUID `db:"scheme_id"`
	CriterionID uuid.UUID `db:"criterion_id"`
}
