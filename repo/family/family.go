package family

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/lftzzzzfeng/fasms/db"
	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/pkg/errors"
)

const DummyAddress = "dummy address"

type Family interface {
	Create(ctx context.Context, family *domain.Family) error
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Family, error)
}

type FamilyRepo struct {
	db db.Execer
}

// New creates an new instance of the repository.
func New(dbExecer db.Execer) Family {
	return &FamilyRepo{
		db: dbExecer,
	}
}

// Create inserts an entry into the database table.
func (r *FamilyRepo) Create(ctx context.Context, data *domain.Family) error {

	sql := `
		INSERT INTO fasms.families (id, address)
		VALUES (:id, :address)
	`

	_, err := r.db.NamedExecContext(ctx, sql, data)
	if err != nil {
		return errors.Wrap(err, "family repository create: insert failed")
	}

	return nil
}

func (r *FamilyRepo) GetByID(ctx context.Context, id uuid.UUID) (*domain.Family, error) {
	query := `
		SELECT id,
			address
		FROM fasms.families
		WHERE id := id
		LIMIT 1
	`

	var family *domain.Family
	err := r.db.QueryRowxContext(ctx, query, id).Scan(family)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrapf(err, "familyrepo: family not found.")
		}

		return nil, errors.Wrapf(err, "familyrepo: get family failed.")
	}

	return family, nil
}
