package response

import "github.com/google/uuid"

type GetEligibleScheme struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}
