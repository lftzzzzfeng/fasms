package domain

import "github.com/google/uuid"

type ApplicantCommon struct {
	ID               uuid.UUID `db:"id"`
	Name             string    `db:"name"`
	Sex              string    `db:"sex"`
	IC               string    `db:"ic"`
	Relationship     string    `db:"relationship"`
	EmploymentStatus string    `db:"employment_status"`
}

type Applicant struct {
	*ApplicantCommon
	FamilyID uuid.UUID `db:"family_id"`
}

type ApplicantFamily struct {
	*ApplicantCommon
	Family string `db:"address"`
}
