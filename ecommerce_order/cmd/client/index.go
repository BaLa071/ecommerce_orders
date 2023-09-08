package main

import (
	"ecommerce_order/cmd/client/grpcclient"
	"ecommerce_order/cmd/client/route"

	"github.com/gin-gonic/gin"
)

func main() {
	_, conn := grpcclient.GetGrpcClientInstance()
	defer conn.Close()

	r := gin.Default()
	route.AppRoutes(r)
	r.Run(":8000")

}

// import (
// 	"context"
// 	"fmt"

// 	// "fmt"

// 	"log"
// 	"net/http"

// 	// models "ecommerce_order/order_dal/models"
// 	// "ecommerce_order/order_dal/models"
// 	"ecommerce_order/order_config/constants"
// 	"ecommerce_order/order_controller/controller"
// 	pb "ecommerce_order/order_proto"

// 	"github.com/gin-gonic/gin"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"google.golang.org/grpc"
// )

// var (
// 	mongoclient *mongo.Client
// 	ctx         context.Context
// 	server      *gin.Engine
// )

// func main() {
// 	// r := gin.Default()
// 	r := gin.Default()
// 	conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
// 	if err != nil {
// 		log.Fatalf("Failed to connect: %v", err)
// 	}
// 	defer conn.Close()

// 	client := pb.NewOrderServiceClient(conn)

// 	r.POST("/createorder", func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
// 		fmt.Println(err1)

// 		var request pb.CustomerOrder
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		request.CustomerId = result
// 		response, err := client.CreateOrder(c.Request.Context(), &request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})

// 	r.POST("/updateorder/:customerid", func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
// 		fmt.Println(err1)

// 		var request pb.UpdateOrderRequest
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		request.Customer_ID = result
// 		response, err := client.UpdateOrderDetails(c.Request.Context(), &request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})

// 	r.POST("/addorderdetails", func(c *gin.Context) {
// 		token := c.GetHeader("Authorization")
// 		result, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
// 		fmt.Println(err1)
// 		var request pb.UpdateOrderRequest
// 		if err := c.ShouldBindJSON(&request); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		request.Customer_ID = result
// 		response, err := client.AddOrderDetails(c.Request.Context(), &request)
// 		if err != nil {
// 			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 			return
// 		}
// 		c.JSON(http.StatusOK, gin.H{"value": response})
// 	})

// 	r.GET("/deletecustomer", func(c *gin.Context) {
// 		var user pb.RemoveOrderRequest
// 		token := c.GetHeader("Authorization")
// 		result1, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
// 		fmt.Println(err1)
// 		if err := c.ShouldBindJSON(&user); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		_, err := client.RemoveOrderCustomer(c.Request.Context(), &pb.RemoveOrderRequest{CustomerId: result1})
// 		if err != nil {
// 			c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 		}
// 		c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

// 	})

// 	r.GET("/getbyid", func(c *gin.Context) {
// 		var res pb.GetOrderRequest
// 		token := c.GetHeader("Authorization")
// 		result1, err1 := controller.ExtractCustomerID(token, constants.SecretKey)
// 		fmt.Println(err1)
// 		if err := c.ShouldBindJSON(&res); err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 			return
// 		}
// 		fmt.Println(res.CustomerId)
// 		result, err := client.GetOrderDetails(c.Request.Context(), &pb.GetOrderRequest{CustomerId: result1})
// 		if err != nil {
// 			c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 		}
// 		c.JSON(http.StatusCreated, gin.H{"status": "success", "data": result})
// 	})

// 	// r.POST("/getbyid", func(c *gin.Context) {
// 	// 	var res pb.GetOrderRequest
// 	// 	result, err := client.GetOrderDetails(c.Request.Context(), &pb.GetOrderRequest{CustomerId: res.CustomerId})
// 	// 	if err != nil {
// 	// 		c.JSON(http.StatusBadGateway, gin.H{"status": "fail", "message": err.Error()})
// 	// 	}
// 	// 	c.JSON(http.StatusCreated, gin.H{"status": "success", "data": result})
// 	// })
// 	r.Run(":8080")
// 	// _, err1 := client.CreateOrder(context.Background(), &pb.CustomerOrder{
// 	//  CustomerId:    "11",
// 	//  PaymentId:     "your_payment_id",
// 	//  PaymentStatus: "your_payment_status",
// 	//  Status:        "your_status",
// 	//  Currency:      "your_currency",
// 	//  Items:[]*pb.Items{
// 	//      {
// 	//          Sku:         "SKU002",
// 	//          Quantity:    5,

// 	//          // Discount:    5.67,  // Your discount value
// 	//          // PreTaxTotal: 18.01, // Your pre-tax total value
// 	//          // Tax:         1.23,  // Your tax value
// 	//          // Total:       19.24, // Your total value
// 	//      },
// 	//      // {
// 	//      //  Sku:         "SKU001",
// 	//      //  Quantity:    3,
// 	//      // },
// 	//      // Add more items if needed
// 	//  },
// 	//  Shipping: []*pb.Shipping{
// 	//      {
// 	//         Address:[]*pb.Address{
// 	//          {
// 	//                 Street1: "Anna Nagar",
// 	//              Street2: "Gandhi nagar",
// 	//              City: "Chennai",
// 	//              State: "TamilNadu",
// 	//              Country: "India",
// 	//              Zip: "56456",
// 	//          },

// 	//         },
// 	//         Origin: []*pb.Origin{
// 	//          {
// 	//              Street1: "Anna Nagar",
// 	//              Street2: "Gandhi nagar",
// 	//              City: "Chennai",
// 	//              State: "TamilNadu",
// 	//              Country: "India",
// 	//              Zip: "56456",
// 	//          },
// 	//         },
// 	//      },

// 	//  },
// 	//  Carrier:  "your_carrier",
// 	//  Tracking: "your_tracking",
// 	// })
// 	// r.POST("/create", func(c *gin.Context) {
// 	//  var request pb.CustomerOrder
// 	//  if err := c.ShouldBindJSON(&request); err != nil {
// 	//      c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 	//      return
// 	//  }
// 	//  response, err := client.CreateOrder(c.Request.Context(), &request)
// 	//  if err != nil {
// 	//      c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 	//      return
// 	//  }
// 	//  c.JSON(http.StatusOK, gin.H{"value": response})
// 	// })

// 	// _, err1 := client.UpdateOrderDetails(context.Background(), &pb.UpdateOrderRequest{
// 	//         Customer_ID:    "11",
// 	//          Sku:         "SKU002",
// 	//          Quantity:    7,

// 	//      // Add more items if needed
// 	//  })

// 	//  _, err1 := client.AddOrderDetails(context.Background(), &pb.UpdateOrderRequest{
// 	//      Customer_ID:   "97",
// 	//       Sku:         "SKU002",
// 	//       Quantity:    7.0,

// 	//   // Add more items if needed
// 	//  })

// 	//  // Shipping: []*pb.Shipping{
// 	// {
// 	//    Address:[]*pb.Address{
// 	//  {
// 	//         Street1: "Anna Nagar",
// 	//      Street2: "Gandhi nagar",
// 	//      City: "Chennai",
// 	//      State: "TamilNadu",
// 	//      Country: "India",
// 	//      Zip: "56456",
// 	//  },

// 	//    },
// 	//    Origin: []*pb.Origin{
// 	//  {
// 	//      Street1: "Anna Nagar",
// 	//      Street2: "Gandhi nagar",
// 	//      City: "Chennai",
// 	//      State: "TamilNadu",
// 	//      Country: "India",
// 	//      Zip: "56456",
// 	//  },
// 	//    },
// 	// },

// 	// },
// 	//)

// 	// _, err1 := client.RemoveOrderCustomer(context.Background(), &pb.RemoveOrderRequest{
// 	//  CustomerId:345,
// 	// })

// 	//  _, err1 := client.GetOrderDetails(context.Background(), &pb.GetOrderRequest{
// 	//  CustomerId:76,
// 	// })
// 	// if err1 != nil {
// 	//  log.Fatalf("Failed to call SayHello: %v", err1)
// 	// }
// 	// r.Run(":8080")
// }
