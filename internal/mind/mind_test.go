package mind_test

import (
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	model "github.com/svtter/mind/internal"
	apperr "github.com/svtter/mind/internal/errors"
	"github.com/svtter/mind/internal/mind"
	"github.com/svtter/mind/internal/mock"
	"github.com/svtter/mind/internal/mock/mockdb"
)

func TestView(t *testing.T) {
	type args struct {
		c  *gin.Context
		id int
	}
	cases := []struct {
		name     string
		args     args
		wantData *model.Mind
		wantErr  error
		mdb      *mockdb.Mind
	}{
		{
			name:    "Fail on RBAC",
			args:    args{id: 5},
			wantErr: apperr.Forbidden,
		},
		{
			name: "Success",
			args: args{id: 1},
			wantData: &model.Mind{
				Base: model.Base{
					ID:        1,
					CreatedAt: mock.TestTime(2000),
					UpdatedAt: mock.TestTime(2000),
				},
				Title:   "John",
				Content: "Doe",
			},
			mdb: &mockdb.Mind{
				ViewFn: func(ctx context.Context, id int) (*model.Mind, error) {
					if id == 1 {
						return &model.Mind{
							Base: model.Base{
								ID:        1,
								CreatedAt: mock.TestTime(2000),
								UpdatedAt: mock.TestTime(2000),
							},
							Title:   "John",
							Content: "Doe",
						}, nil
					}
					return nil, nil
				}},
		},
	}
	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			s := mind.New(tt.mdb)
			usr, err := s.View(tt.args.c, tt.args.id)
			assert.Equal(t, tt.wantData, usr)
			assert.Equal(t, tt.wantErr, err)
		})
	}
}
