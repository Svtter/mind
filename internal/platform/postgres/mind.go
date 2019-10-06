package pgsql

import (
	"context"

	"github.com/go-pg/pg"
	model "github.com/svtter/mind/internal"
	apperr "github.com/svtter/mind/internal/errors"
	"go.uber.org/zap"
)

// NewMindDB returns a new UserDB instance
func NewMindDB(c *pg.DB, l *zap.Logger) *MindDB {
	return &MindDB{c, l}
}

// MindDB represents the client for Mind table.
type MindDB struct {
	cl  *pg.DB
	log *zap.Logger
}

// View returns
func (m *MindDB) View(c context.Context, id int) (*model.Mind, error) {
	base := model.Base{ID: id}
	mind := &model.Mind{Base: base}
	err := m.cl.Select(mind)
	if err != nil {
		m.log.Warn("MindDB Error", zap.Error(err))
		return nil, apperr.NotFound
	}
	return mind, nil
}
