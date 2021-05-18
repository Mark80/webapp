package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"webapp/api"
	"webapp/dal"
	"webapp/services"
)

func main() {

	dsn := "host=localhost user=postgres password=password dbname=test port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return
	}

	ctx := context.Background()
	dao := dal.OrderDao{DB: db}

	err = dao.Migrate(ctx)
	if err != nil {
		log.Fatalln(err)
		return
	}

	engine := gin.Default()
	orderService := services.OrderService{
		Repository: dao,
	}

	api.NewRoutes(engine,&orderService)

	err = engine.Run()
	if err != nil {
		log.Fatalln(err)
		return
	}

}
