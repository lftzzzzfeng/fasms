package domain

import "github.com/google/uuid"

// Scheme represents a log entry for a successfull dispatch.
type Scheme struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}

type SchemeInfo struct {
	SchemeID      uuid.UUID `db:"scheme_id"`
	Name          string    `db:"name"`
	Description   string    `db:"description"`
	CriID         uuid.UUID `db:"c_id"`
	Criterion     string    `db:"criterion"`
	CriDetail     string    `db:"c_detail"`
	BnftID        uuid.UUID `db:"b_id"`
	Benefit       string    `db:"benefit"`
	BenefitDetail string    `db:"b_detail"`
}
