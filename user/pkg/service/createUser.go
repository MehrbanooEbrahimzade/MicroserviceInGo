package service

import (
	"context"
	pb "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
	"log"
)

func (srv *userService) CreateService(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	userId, err := srv.repo.CreateRepo(ctx, req)
	if err != nil {
		log.Println(err)
	}
	return userId, nil
}
