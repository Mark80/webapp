package dal

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"webapp/business_models"
)

type OrderRepository interface {
	Migrate(ctx context.Context) error
	Save(ctx context.Context, o *business_models.Order) (uint, error)
    GetAll(ctx context.Context) ([]business_models.Order, error)
}
type OrderDao struct {
	DB *gorm.DB
}

type order struct {
	gorm.Model
	OrderID string
	User    string
	Item    string
}

func (b OrderDao) Migrate(ctx context.Context) error {
	return b.DB.WithContext(ctx).AutoMigrate(&order{})
}

func (b OrderDao) Save(ctx context.Context, o *business_models.Order) (uint, error) {
	orderDb := &order{
		OrderID: o.OrderID,
		User:    o.User,
		Item:    o.Item,
	}

	res := b.DB.WithContext(ctx).Save(orderDb)
	if res.Error != nil {
		log.Errorf("failed to insert order, %v\n", res.Error)
		return 0, fmt.Errorf("failed to insert order, %w", res.Error)
	}
	return orderDb.ID, nil

}

func (b OrderDao) GetAll(ctx context.Context) ([]business_models.Order, error) {

	var orders []order
	res := b.DB.WithContext(ctx).Find(&orders)
	if res.Error != nil {
		log.Errorf("failed to retrive orders, %v\n", res.Error)
		return nil, fmt.Errorf("failed to retrive orders, %w", res.Error)
	}

	var bOrders []business_models.Order
	for _, o := range orders {

		bo := business_models.Order{
			OrderID: o.OrderID,
			User:    o.User,
			Item:    o.Item,
		}
		bOrders = append(bOrders, bo)
	}

	return bOrders, nil

}
