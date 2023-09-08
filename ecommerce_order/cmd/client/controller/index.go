package controller

import (
	grpcclient "ecommerce_order/cmd/client/grpcclient"
	"ecommerce_order/order_config/constants"
	"ecommerce_order/order_controller/controller"
	pb "ecommerce_order/order_proto"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerCreateOrder(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()
	token := c.GetHeader("Authorization")
	result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
	fmt.Println(err1)

	var request pb.CustomerOrder
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.CustomerId = result
	response, err := grpcClient.CreateOrder(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": response.CustomerId})

}

func HandlerUpdateOrder(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()
	token := c.GetHeader("Authorization")
	result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
	fmt.Println(err1)

	var request pb.UpdateOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.Customer_ID = result
	response, err := grpcClient.UpdateOrderDetails(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": response})

}

func HandlerAddOrder(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()
	token := c.GetHeader("Authorization")
	result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
	fmt.Println(err1)

	var request pb.UpdateOrderRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	request.Customer_ID = result
	response, err := grpcClient.UpdateOrderDetails(c.Request.Context(), &request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"value": response})

}

func HandlerDeleteOrder(c *gin.Context) {
	grpcClient, _ := grpcclient.GetGrpcClientInstance()
	token := c.GetHeader("Authorization")
	result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
	fmt.Println(err1)

	var user pb.RemoveOrderRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := grpcClient.RemoveOrderCustomer(c.Request.Context(), &pb.RemoveOrderRequest{CustomerId: result})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

}

func HandlerGetOrderById(c *gin.Context) {
	var res pb.GetOrderRequest
	grpcClient, _ := grpcclient.GetGrpcClientInstance()
	token := c.GetHeader("Authorization")
	result1, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
	fmt.Println(err1)
	if err := c.ShouldBindJSON(&res); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println(res.CustomerId)
	result, err := grpcClient.GetOrderDetails(c.Request.Context(), &pb.GetOrderRequest{CustomerId: result1})
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
	}
	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": result})
}
