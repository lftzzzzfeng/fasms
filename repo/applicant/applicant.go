package applicant

import (
	"context"
	"fmt"

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
		fmt.Println("err", err)
		return errors.Wrap(err, "applicant repository Create: insert failed")
	}

	return nil
}

// GetAll return all applicants
func (r *ApplicantRepo) GetAll(ctx context.Context) ([]*domain.Applicant, error) {
	return nil, nil
}
