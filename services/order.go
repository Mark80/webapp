package services

import (
	"context"
	"webapp/business_models"
	"webapp/dal"
)

type OrderService struct {
	Repository dal.OrderRepository
}

func (service OrderService)Save(ctx context.Context, o *business_models.Order) (uint, error)  {
	return service.Repository.Save(ctx, o)
}
