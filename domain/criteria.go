package domain

import "github.com/google/uuid"

var criteriaMap = map[string]string{
	"employ_status":    "umemployed",
	"has_child":        "primary",
	"residence_status": "citizen",
}

type Criteria struct {
	ID     uuid.UUID `db:"id"`
	Name   string    `db:"name"`
	Detail string    `db:"detail"`
}
