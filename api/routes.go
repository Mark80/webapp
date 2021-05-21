package api

import (
	"github.com/gin-gonic/gin"
	"webapp/services"
)

type Routes struct {
	engine       *gin.Engine
	orderService *services.OrderService
}

func NewRoutes(engine *gin.Engine, orderService *services.OrderService) {
	engine.POST("api/orders", SaveOrder(orderService))
	engine.GET("api/orders", GetAllOrders(orderService))
	engine.GET("api/orders/:id", GetOrderById(orderService))
}

func GetOrderById(orderService *services.OrderService) func(context *gin.Context) {
	return func(context *gin.Context) {
		id, existParams := context.Params.Get("id")
		if !existParams {
			context.JSON(500, gin.H{"error": "generic error"})
		}
		order, err := orderService.GetByID(context, id)
		if err != nil {
			context.JSON(500, gin.H{"error": "generic error"})
		} else {
			context.JSON(200, gin.H{"orders": order})
		}
	}
}

func GetAllOrders(orderService *services.OrderService) func(context *gin.Context) {
	return func(context *gin.Context) {
		orders, err := orderService.GetAll(context)
		if err != nil {
			context.JSON(500, gin.H{"error": "generic error"})
		} else {
			context.JSON(200, gin.H{"orders": orders})
		}
	}
}

func SaveOrder(orderService *services.OrderService) func(context *gin.Context) {
	return func(context *gin.Context) {
		var order services.Order
		err := context.BindJSON(&order)
		if err != nil {
			context.JSON(400, gin.H{"error": "Parsing json error"})
		} else {
			save, err := orderService.Save(context, &order)
			if err != nil {
				context.JSON(500, gin.H{"error": "generic error"})
			} else {
				context.JSON(200, gin.H{"id": save})
			}

		}
	}
}
