package domain

import "github.com/google/uuid"

type Criteria struct {
	ID     uuid.UUID `db:"id"`
	Name   string    `db:"name"`
	Detail string    `db:"detail"`
}
