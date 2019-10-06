package pgsql

import (
	"context"
	"log"
	"time"

	"github.com/go-pg/pg"
)

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, event *pg.QueryEvent) (context.Context, error) {
	query, err := event.FormattedQuery()
	checkErr(err)
	log.Printf("%s | %s", time.Since(event.StartTime), query)
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, event *pg.QueryEvent) error {
	return nil
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
