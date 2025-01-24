package domain

import "github.com/google/uuid"

type Family struct {
	ID      uuid.UUID `db:"id"`
	Address string    `db:"address"`
}
