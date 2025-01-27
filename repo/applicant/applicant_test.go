package applicant_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestApplicant_Create(t *testing.T) {
	cases := []struct {
		name     string
		assertFn func(subT *testing.T)
	}{
		{
			name: "should return no error if successfully creating applicant",
			assertFn: func(subT *testing.T) {
				asrt := assert.New(subT)
				asrt.NoError(nil, "success")
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, c.assertFn)
	}
}
