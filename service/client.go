package service

import (
	"demo_gateway/pb"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceClient struct {
	Person pb.PersonManagementClient
}

func NewServiceClient() *serviceClient {
	clientConn, err := grpc.Dial(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Print("fail to connect to person service")
	}

	return &serviceClient{
		Person: pb.NewPersonManagementClient(clientConn),
	}
}
