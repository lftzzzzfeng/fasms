package scheme

import (
	"context"

	schemerepo "github.com/lftzzzzfeng/fasms/repo/scheme"
)

type Scheme struct {
	SchemeRepo schemerepo.Scheme
}

func New(schemeRepo schemerepo.Scheme) *Scheme {
	return &Scheme{
		SchemeRepo: schemeRepo,
	}
}

func (s *Scheme) GetAllSchemes(ctx context.Context) {

}
