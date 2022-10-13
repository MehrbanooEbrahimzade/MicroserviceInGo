package main

import (
	"context"
	"fmt"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/cmd"
	pb "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
	"log"
	"net"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	factory, err := cmd.NewFactory(ctx)
	lis, err := net.Listen("tcp", factory.Config.Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pb.RegisterUserServiceServer(factory.GrpcServer, factory.GrpcHandler)
	log.Printf("server listening at %v", lis.Addr())

	if err != nil {
		log.Println(err)
	}
	fmt.Printf("Server succesfully started on port :%v", factory.Config.Port)

	ch := make(chan os.Signal)
	signal.Notify(ch, os.Interrupt)
	<-ch
	fmt.Println("\nStopping the server...")
	factory.GrpcServer.Stop()
	lis.Close()
	fmt.Println("Closing MongoDB connection")
	factory.Db.Client.Disconnect(ctx)
	fmt.Println("Done.")
}
