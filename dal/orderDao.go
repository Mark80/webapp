package dal

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"webapp/services"
)

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

func (b OrderDao) Save(ctx context.Context, o services.Order) (uint, error) {
	orderDb := order{
		OrderID: o.OrderID,
		User:    o.User,
		Item:    o.Item,
	}

	res := b.DB.WithContext(ctx).Save(&orderDb)
	if res.Error != nil {
		log.Errorf("failed to insert order, %v\n", res.Error)
		return 0, fmt.Errorf("failed to insert order, %w", res.Error)
	}
	return orderDb.ID, nil

}

func (b OrderDao) GetAll(ctx context.Context) ([]services.Order, error) {

	var orders []order
	res := b.DB.WithContext(ctx).Find(&orders)
	if res.Error != nil {
		log.Errorf("failed to retrive orders, %v\n", res.Error)
		return nil, fmt.Errorf("failed to retrive orders, %w", res.Error)
	}

	var bOrders []services.Order
	for _, o := range orders {
		bOrders = append(bOrders, convertFrom(o))
	}
	return bOrders, nil
}

func (b OrderDao) GetByID(ctx context.Context, id string) (services.Order, error) {

	var o order
	res := b.DB.WithContext(ctx).First(&o, id)
	if res.Error != nil {
		log.Errorf("failed to retrive orders, %v\n", res.Error)
		return services.Order{}, fmt.Errorf("failed to retrive orders, %w", res.Error)
	}

	result := convertFrom(o)
	return result, nil

}

func convertFrom(o order) services.Order {
	return services.Order{
		OrderID: o.OrderID,
		User:    o.User,
		Item:    o.Item,
	}
}
