package response

import "github.com/google/uuid"

type GetAllApplicant struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
