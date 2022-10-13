package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Mongodb struct {
	Client *mongo.Client
}

func Initialdb(port string) *Mongodb {
	//lis, err := net.Listen("tcp", port)
	//if err != nil {
	//	log.Fatalf("failed to listen: %v", err)
	//}
	//var opts []grpc.ServerOption
	//s := grpc.NewServer(opts...)
	//gRPCInGo.RegisterUserServiceServer(s, &UserServiceServer{})
	//log.Printf("mongodb listening at %v", lis.Addr())
	//ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	////mongodb|
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.NewClient(options.Client().ApplyURI(port))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	}

	return &Mongodb{
		Client: client,
	}
}
