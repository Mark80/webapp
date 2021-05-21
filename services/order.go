package services

import (
	"context"
	"fmt"
)

type OrderNotFound struct {
	Id string
}

func (err *OrderNotFound) Error() string {
	return fmt.Sprintf("Order with Id %s not found", err.Id)
}

type OrderRepository interface {
	Migrate(ctx context.Context) error
	Save(ctx context.Context, o *Order) (uint, error)
	GetAll(ctx context.Context) ([]Order, error)
	GetByID(ctx context.Context, id  string) (*Order, error)
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

func (service OrderService)GetByID(ctx context.Context,id string)  (*Order,error) {
	return service.Repository.GetByID(ctx, id)
}
