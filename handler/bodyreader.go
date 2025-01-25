package handler

import (
	"encoding/json"
	"io"

	"github.com/pkg/errors"
)

// The maximum size of the request body.
const MaxBodySize = 1024 * 4

// page size
const PageSize = 10

// ReadJSON reads the JSON encoded value from the provider reader and stores
// it in the values pointed to by v.
func (router *Router) readJSON(r io.Reader, v any) error {
	decoder := json.NewDecoder(io.LimitReader(r, MaxBodySize))

	err := decoder.Decode(v)
	if err != nil {
		return errors.Wrap(err, "ReadJSON: JSON read failed")
	}

	return nil
}
