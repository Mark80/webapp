package api

import (
	"github.com/gin-gonic/gin"
	"webapp/business_models"
	"webapp/services"
)

type Routes struct {
	engine       *gin.Engine
	orderService *services.OrderService
}

func NewRoutes(engine *gin.Engine, orderService *services.OrderService) {

	engine.POST("api/orders", func(context *gin.Context) {
		var order business_models.Order
		err := context.BindJSON(&order)
		if err != nil {
			context.JSON(400, gin.H{"error": "Parsing json error"})
		} else {
			save, err := orderService.Save(context, &order)
			if err != nil {
				context.JSON(500, gin.H{"error": "generic error"})
			}else {
				context.JSON(200, gin.H{"id": save})
			}

		}
	})

	engine.GET("api/orders", func(context *gin.Context) {
		orders, err := orderService.GetAll(context)
		if err != nil {
			context.JSON(500, gin.H{"error": "generic error"})
		}else{
			context.JSON(200, gin.H{"orders": orders})
		}
	})

}
