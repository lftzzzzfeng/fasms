package domain

import "github.com/google/uuid"

type Benefit struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
