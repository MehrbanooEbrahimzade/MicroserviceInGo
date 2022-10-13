package handler

import (
	"context"
	pb "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/service"
	"log"
)

type UserHandler struct {
	srv service.Iservice
	pb.UserServiceServer
}

func NewHandler(srv service.Iservice) (*UserHandler, error) {
	return &UserHandler{
		srv: srv,
	}, nil
}

func (h *UserHandler) Create(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	res, err := h.srv.CreateService(ctx, req)
	if err != nil {
		log.Println(err)
	}
	return res, err
}

func (h *UserHandler) GetAll(req *pb.ReadAllReq, stream pb.UserService_ReadAllServer) error {
	err := h.srv.ReadAllService(req, stream)
	if err != nil {
		log.Println(err)
	}
	return err
}
