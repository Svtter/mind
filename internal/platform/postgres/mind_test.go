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
	userDB := pgsql.NewUserDB(c, l)
	cases := []struct {
		name string
		fn   func(*testing.T, *pgsql.UserDB, *pg.DB)
	}{
		{
			name: "view",
			fn:   testDBView,
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			tt.fn(t, userDB, c)
		})
	}

}

// testDBView
func testDBView(t *testing.T, db *pgsql.UserDB, c *pg.DB) {
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
		{
			name: "Success",
			id:   2,
			wantData: &model.Mind{
				Title:   "test",
				Content: "test",
				User: &model.User{
					FirstName: "John",
				},
			},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			user, err := db.View(nil, tt.id)
			assert.Equal(t, tt.wantErr, err != nil)
			if tt.wantData != nil {
				tt.wantData.CreatedAt = user.CreatedAt
				tt.wantData.UpdatedAt = user.UpdatedAt
				assert.Equal(t, tt.wantData, user)
			}
		})
	}
}
