package grpcclient

import (
	"log"
	"sync"
	pb "ecommerce_order/order_proto"
	"google.golang.org/grpc"
)

var once sync.Once

type GrpcClient pb.OrderServiceClient

var (
	instance GrpcClient
)

func GetGrpcClientInstance() (GrpcClient, *grpc.ClientConn) {
	var conn *grpc.ClientConn
	once.Do(func() { // <-- atomic, does not allow repeating
		conn, err := grpc.Dial("localhost:5002", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Failed to connect: %v", err)
		}
		//defer conn.Close()

		instance = pb.NewOrderServiceClient(conn)
	})

	return instance, conn
}
