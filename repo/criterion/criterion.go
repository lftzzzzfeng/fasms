package criterion

import (
	"context"
	"database/sql"
	"strings"

	"github.com/google/uuid"
	"github.com/lftzzzzfeng/fasms/db"
	"github.com/pkg/errors"
)

const (
	CRITERION_CITIZEN        = "citizen"
	CRITERION_PRIMARY_SCHOOL = "primary"
)

type Criterion interface {
	GetIdsByDetails(ctx context.Context, vals []string) ([]uuid.UUID, error)
}

type CriteriaRepo struct {
	db db.Execer
}

// New creates an new instance of the repository.
func New(dbExecer db.Execer) Criterion {
	return &CriteriaRepo{
		db: dbExecer,
	}
}

// GetAll return all applicants
func (r *CriteriaRepo) GetIdsByDetails(ctx context.Context, vals []string) ([]uuid.UUID, error) {
	query := `
		SELECT id
		FROM fasms.criteria 
		WHERE detail = ANY($1)
	`

	anyStr := "{" + strings.Join(vals, ",") + "}"
	rows, err := r.db.QueryxContext(ctx, query, anyStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "criterionrepo: criterion id not found from db.")
		}

		return nil, errors.Wrap(err, "criterionrepo: get criterion id from db failed.")
	}

	var ids []uuid.UUID
	for rows.Next() {
		var id uuid.UUID

		if err = rows.Scan(&id); err != nil {
			return nil, errors.Wrap(err, "criterionrepo: scan criterion id failed")
		}

		ids = append(ids, id)
	}

	return ids, nil
}
