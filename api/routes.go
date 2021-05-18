package api

import (
	"github.com/gin-gonic/gin"
	"webapp/business_models"
	"webapp/services"
)

type Routes struct {
	engine *gin.Engine
	orderService *services.OrderService
}

func NewRoutes(engine *gin.Engine, orderService *services.OrderService) {

	engine.POST("orders", func(context *gin.Context) {
		var order business_models.Order
		err := context.BindJSON(&order)
		if err != nil {
			context.JSON(400, gin.H{"error": "Parsing json error"})
		}else {
			orderService.Save(context, &order)
		}
	})


}



