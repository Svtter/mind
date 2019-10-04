package mockdb

import (
	"context"

	model "github.com/svtter/mind/internal"
)

// Mind database mock
type Mind struct {
	ViewFn func(context.Context, int) (*model.Mind, error)
}

// View mock
func (m *Mind) View(c context.Context, i int) (*model.Mind, error) {
	return m.ViewFn(c, i)
}
