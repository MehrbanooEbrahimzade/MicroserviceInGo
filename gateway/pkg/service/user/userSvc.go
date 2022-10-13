package userSvc

import (
	"context"
	users "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/proto/user"
	userRepo "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/repository/user"
)

type UserSvc struct {
	repo *userRepo.UserRepo
	ctx  context.Context
}

func NewUserSvc(ctx context.Context, client *userRepo.UserRepo) *UserSvc {
	return &UserSvc{
		repo: client,
		ctx:  ctx,
	}
}

func (u *UserSvc) GetAllUsers(req *users.ReadAllReq) (res *users.ReadAllRes, err error) {
	res, err = u.repo.GetUsers(u.ctx, req)
	if err != nil {
		return nil, err
	}

	return
}
func (u *UserSvc) CreateUser(req *users.CreateUserReq) (res *users.CreateUserRes, err error) {
	res, err = u.repo.CreateUser(u.ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
