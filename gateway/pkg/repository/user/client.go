package userRepo

import (
	users "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/proto/user"
	"log"

	"google.golang.org/grpc"
)

func NewUserRepoClient(port string) *users.UserServiceClient {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	c := users.NewUserServiceClient(conn)
	return &c
}
