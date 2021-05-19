package services

import (
	"context"
)

type OrderRepository interface {
	Migrate(ctx context.Context) error
	Save(ctx context.Context, o *Order) (uint, error)
	GetAll(ctx context.Context) ([]Order, error)
}

type OrderService struct {
	Repository OrderRepository
}

func (service OrderService)Save(ctx context.Context, o *Order) (uint, error)  {
	return service.Repository.Save(ctx, o)
}

func (service OrderService)GetAll(ctx context.Context) ([]Order, error)  {
	return service.Repository.GetAll(ctx)
}
