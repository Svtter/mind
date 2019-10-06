package pgsql

import (
	"time"

	"github.com/go-pg/pg"
	// DB adapter
	_ "github.com/lib/pq"
	model "github.com/svtter/mind/internal"

	"github.com/svtter/mind/cmd/api/config"
)

const notDeleted = "deleted_at is null"

// New creates new database connection to a postgres database
// Function panics if it can't connect to database
func New(cfg *config.Database) (*pg.DB, error) {
	u, err := pg.ParseURL(cfg.PSN)
	if err != nil {
		return nil, err
	}
	db := pg.Connect(u).WithTimeout(time.Second * 5)
	_, err = db.Exec("SELECT 1")
	if err != nil {
		return nil, err
	}
	if cfg.Log {
		db.AddQueryHook(dbLogger{})
	}
	if cfg.CreateSchema {
		createSchema(db, &model.Company{}, &model.Location{}, &model.Role{}, &model.User{}, &model.Mind{})
	}
	return db, nil
}

func createSchema(db *pg.DB, models ...interface{}) {
	for _, model := range models {
		checkErr(db.CreateTable(model, nil))
	}
}
