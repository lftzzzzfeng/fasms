package response

import "github.com/google/uuid"

type GetAllSchemes struct {
	ID          uuid.UUID    `json:"id"`
	Name        string       `json:"name"`
	Description string       `json:"descritpion"`
	Criteria    []*Criterion `json:"criteria"`
	Benefits    []*Benefit   `json:"benefits"`
}

type Criterion struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Detail string    `json:"detail"`
}

type Benefit struct {
	ID     uuid.UUID `json:"id"`
	Name   string    `json:"name"`
	Detail string    `json:"detail"`
}
