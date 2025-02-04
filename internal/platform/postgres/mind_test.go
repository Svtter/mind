package pgsql_test

import (
	"testing"

	"github.com/go-pg/pg"
	"github.com/stretchr/testify/assert"
	model "github.com/svtter/mind/internal"
	pgsql "github.com/svtter/mind/internal/platform/postgres"
	"go.uber.org/zap"
)

func testMindDB(t *testing.T, c *pg.DB, l *zap.Logger) {
	mindDB := pgsql.NewMindDB(c, l)
	cases := []struct {
		name string
		fn   func(*testing.T, *pgsql.MindDB, *pg.DB)
	}{
		{
			name: "view",
			fn:   testDBView,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn(t, mindDB, c)
		})
	}

}

// testDBView
func testDBView(t *testing.T, db *pgsql.MindDB, c *pg.DB) {
	cases := []struct {
		name     string
		wantErr  bool
		id       int
		wantData *model.Mind
	}{
		{
			name:    "Mind does not exist",
			wantErr: true,
			id:      1000,
		},
		// {
		// 	name: "Success",
		// 	id:   1,
		// 	wantData: &model.Mind{
		// 		Base:    model.Base{ID: 1},
		// 		Title:   "test",
		// 		Content: "test",
		// 		User: &model.User{
		// 			Base: model.Base{ID: 1},
		// 		},
		// 	},
		// },
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			mind, err := db.View(nil, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData != nil {
				tt.wantData.CreatedAt = mind.CreatedAt
				tt.wantData.UpdatedAt = mind.UpdatedAt
				assert.Equal(t, tt.wantData, mind)
			}
		})
	}
}
