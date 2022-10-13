package service

import (
	pb "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
	"log"
)

func (srv *userService) ReadAllService(req *pb.ReadAllReq, stream pb.UserService_ReadAllServer) error {
	err := srv.repo.ReadAllRepo(req, stream)
	if err != nil {
		log.Println(err)
	}
	return err
}
