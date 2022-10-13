package service

import (
	"context"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/IRepo"
	user "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
)

type userService struct {
	repo IRepo.IRepo
}

type Iservice interface {
	CreateService(context.Context, *user.CreateUserReq) (*user.CreateUserRes, error)
	ReadAllService(*user.ReadAllReq, user.UserService_ReadAllServer) error
}

func NewService(repo IRepo.IRepo) (Iservice, error) {
	return &userService{repo: repo}, nil
}
