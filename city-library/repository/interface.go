package repository

import (
	"context"

	"github.com/rbojan2000/city/model"
)

type Repository interface {
	GetBorrow(ctx context.Context, id string) (model.Borrow, error)
	CreateBorrow(ctx context.Context, in model.Borrow) (model.Borrow, error)
	DeleteBorrow(ctx context.Context, id string) error
}
