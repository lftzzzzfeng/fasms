package domain

import "github.com/google/uuid"

type Applicant struct {
	ID               uuid.UUID `db:"id"`
	Name             string    `db:"name"`
	Sex              string    `db:"sex"`
	IC               string    `db:"ic"`
	FamilyID         uuid.UUID `db:"family_id"`
	Relationship     string    `db:"relationship"`
	EmploymentStatus string    `db:"employment_status"`
}
