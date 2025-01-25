package applicant

import (
	"context"
	"database/sql"

	"github.com/lftzzzzfeng/fasms/db"
	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/pkg/errors"
)

type Applicant interface {
	Create(ctx context.Context, applicant *domain.Applicant) error
	GetAll(ctx context.Context) ([]*domain.Applicant, error)
}

type ApplicantRepo struct {
	db db.Execer
}

// New creates an new instance of the repository.
func New(dbExecer db.Execer) Applicant {
	return &ApplicantRepo{
		db: dbExecer,
	}
}

// Create inserts an entry into the database table.
func (r *ApplicantRepo) Create(ctx context.Context, data *domain.Applicant) error {

	sql := `
		INSERT INTO fasms.applicants
			(id, name, sex, ic, family_id, relationship, employment_status)
		VALUES
			(:id, :name, :sex, :ic, :family_id, :relationship, :employment_status)
	`

	_, err := r.db.NamedExecContext(ctx, sql, data)
	if err != nil {
		return errors.Wrap(err, "applicant repository Create: insert failed")
	}

	return nil
}

// GetAll return all applicants
func (r *ApplicantRepo) GetAll(ctx context.Context) ([]*domain.Applicant, error) {

	query := `
		SELECT t.id,
			name,
			sex,
			ic,
			f.address,
			relationship,
			employment_status
		FROM fasms.applicants t
		LEFT JOIN fasms.families f ON f.id = t.family_id
	`

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "applicantrepo: applicant not found from db.")
		}

		return nil, errors.Wrap(err, "applicantrepo: get applicant from db failed.")
	}

	var applicants []*domain.Applicant
	for rows.Next() {
		var applicant *domain.Applicant

		if err = rows.StructScan(applicant); err != nil {
			return nil, errors.Wrap(err, "applicantrepo: scan applicant data failed")
		}

		applicants = append(applicants, applicant)
	}

	return applicants, nil
}
