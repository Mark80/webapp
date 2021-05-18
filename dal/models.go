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
		return 0, fmt.Errorf("failed to store booking interaction, %w", res.Error)
	}
	return orderDb.ID, nil

}


