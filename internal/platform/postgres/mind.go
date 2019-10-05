package pgsql

import (
	"github.com/go-pg/pg"
	"go.uber.org/zap"
)

// NewMindDB returns a new UserDB instance
func NewMindDB(c *pg.DB, l *zap.Logger) *UserDB {
	return &UserDB{c, l}
}
