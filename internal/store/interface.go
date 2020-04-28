package store

import (
	"github.com/bgoldovsky/service-rus-id/internal/domain/valuetypes"
	"github.com/bgoldovsky/service-rus-id/internal/models"
)

type UserRepository interface {
	Get(id valuetypes.UserID) (*models.UserReadModel, error)
	GetList(start, limit *int) ([]*models.UserReadModel, error)
}
