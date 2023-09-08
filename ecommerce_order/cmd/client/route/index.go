package route

import (
	"ecommerce_order/cmd/client/controller"

	"github.com/gin-gonic/gin"
)

func AppRoutes(r *gin.Engine) {

	r.POST("/createorder", controller.HandlerCreateOrder)
	r.POST("/updateorder", controller.HandlerUpdateOrder)
	r.POST("/Addorder", controller.HandlerAddOrder)
	r.GET("/Deleteorder", controller.HandlerDeleteOrder)
	r.GET("/Getorderbyid", controller.HandlerGetOrderById)
}
