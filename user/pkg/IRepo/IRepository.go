package IRepo

import (
	"context"
	pb "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
)

type IRepo interface {
	CreateRepo(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error)
	ReadAllRepo(req *pb.ReadAllReq, stream pb.UserService_ReadAllServer) error
}
