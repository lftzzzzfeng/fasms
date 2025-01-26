package application

import (
	"context"
	"database/sql"

	"github.com/lftzzzzfeng/fasms/db"
	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/pkg/errors"
)

type Application interface {
	Create(ctx context.Context, application *domain.Application) error
	GetAll(ctx context.Context, offset, limit int) ([]*domain.ApplicationInfo, error)
}

type ApplicantionRepo struct {
	db db.Execer
}

// New creates an new instance of the repository.
func New(dbExecer db.Execer) Application {
	return &ApplicantionRepo{
		db: dbExecer,
	}
}

func (a *ApplicantionRepo) Create(ctx context.Context, apl *domain.Application) error {
	sql := `
		INSERT INTO fasms.applications
			(id, applicant_id, scheme_id)
		VALUES
			(:id, :applicant_id, :scheme_id)
	`

	_, err := a.db.NamedExecContext(ctx, sql, apl)
	if err != nil {
		return errors.Wrap(err, "application repository create: insert failed")
	}

	return nil
}

func (a *ApplicantionRepo) GetAll(ctx context.Context, offset, limit int) (
	[]*domain.ApplicationInfo, error) {
	query := `
		SELECT t.id,
			applc.name AS applc_name,
			s.name AS scheme_name
		FROM fasms.applications t
		JOIN fasms.applicants applc ON applc.id = t.applicant_id
		JOIN fasms.schemes s ON s.id = t.scheme_id
		OFFSET $1
		LIMIT $2
	`

	rows, err := a.db.QueryxContext(ctx, query, offset, limit)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "applicationrepo: application not found from db.")
		}

		return nil, errors.Wrap(err, "applicationrepo: get application from db failed.")
	}

	var applications []*domain.ApplicationInfo
	for rows.Next() {
		var application domain.ApplicationInfo

		if err = rows.StructScan(&application); err != nil {
			return nil, errors.Wrap(err, "applicantrepo: scan applicant data failed")
		}

		applications = append(applications, &application)
	}

	return applications, nil
}
