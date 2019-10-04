package mind

import (
	"github.com/gin-gonic/gin"
	model "github.com/svtter/mind/internal"
)

// Service represents the mind application service
type Service struct {
	mdb model.MindDB
}

// New creates new user application service
func New(mdb model.MindDB) *Service {
	return &Service{mdb: mdb}
}

// View returns single user
func (s *Service) View(c *gin.Context, id int) (*model.User, error) {
	return s.mdb.View(c, id)
}
