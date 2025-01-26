package scheme

import (
	"context"
	"database/sql"

	"github.com/lftzzzzfeng/fasms/db"
	"github.com/lftzzzzfeng/fasms/domain"
	"github.com/pkg/errors"
)

type Scheme interface {
	GetAll(ctx context.Context) ([]*domain.SchemeInfo, error)
	GetEligibleSchemesByCritieria(ctx context.Context, criStr string) (
		[]*domain.Scheme, error)
}

type SchemeRepo struct {
	db db.Execer
}

// New creates an new instance of the repository.
func New(dbExecer db.Execer) Scheme {
	return &SchemeRepo{
		db: dbExecer,
	}
}

func (r *SchemeRepo) GetAll(ctx context.Context) ([]*domain.SchemeInfo, error) {
	query := `
		SELECT t.id AS scheme_id,
			t.name,
			t.description,
			c.id AS c_id,
			c.name AS criterion,
			c.detail AS c_detail,
			b.id AS b_id,
			b.name AS benefit,
			b.description AS b_detail
		FROM fasms.schemes t
		LEFT JOIN fasms.scheme_criterion_mapping scm ON scm.scheme_id = t.id
		LEFT JOIN fasms.criteria c ON c.id = scm.criterion_id
		LEFT JOIN fasms.scheme_benefit_mapping sbm ON sbm.scheme_id = t.id
		LEFT JOIN fasms.benefits b ON b.id = sbm.benefit_id
	`

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "schemerepo: scheme not found from db.")
		}

		return nil, errors.Wrap(err, "schemerepo: get scheme from db failed.")
	}

	var schemeInfoArr []*domain.SchemeInfo
	for rows.Next() {
		var schemeInfo domain.SchemeInfo

		if err = rows.StructScan(&schemeInfo); err != nil {
			return nil, errors.Wrap(err, "schemerepo: scan scheme data failed")
		}

		schemeInfoArr = append(schemeInfoArr, &schemeInfo)
	}

	return schemeInfoArr, nil
}

func (r *SchemeRepo) GetEligibleSchemesByCritieria(ctx context.Context,
	criStr string) ([]*domain.Scheme, error) {
	query := `
		SELECT t.id,
			t.name,
			t.description
		FROM fasms.schemes t
		LEFT JOIN fasms.scheme_criterion_mapping scm ON scm.scheme_id = t.id
		LEFT JOIN fasms.criteria c ON c.id = scm.criterion_id
		GROUP BY t.id
		HAVING string_agg(c.id::text,'|') = $1
	`

	rows, err := r.db.QueryxContext(ctx, query, criStr)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errors.Wrap(err, "schemerepo: scheme not found from db.")
		}

		return nil, errors.Wrap(err, "schemerepo: get scheme from db failed.")
	}

	var schemes []*domain.Scheme
	for rows.Next() {
		var scheme domain.Scheme

		if err = rows.StructScan(&scheme); err != nil {
			return nil, errors.Wrap(err, "schemerepo: scan scheme data failed")
		}

		schemes = append(schemes, &scheme)
	}

	return schemes, nil
}
