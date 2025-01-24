package domain

import "github.com/google/uuid"

// DispatchLog represents a log entry for a successfull dispatch.
type DispatchLog struct {
	ID          uuid.UUID `db:"id"`
	Name        string    `db:"name"`
	Description string    `db:"description"`
}
