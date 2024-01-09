package repository

import (
	"context"

	"github.com/rbojan2000/central-library/model"
)

type Repository interface {
	GetUser(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, in model.User) (model.User, error)
	UpdateUser(ctx context.Context, in model.User) (model.User, error)
	DeleteUser(ctx context.Context, id string) error
}
