package userRepo

import (
	"context"
	users "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/proto/user"
)

type UserRepo struct {
	usersClient users.UserServiceClient
}

func NewUsersRepo(client users.UserServiceClient) *UserRepo {
	return &UserRepo{
		usersClient: client,
	}
}

func (u *UserRepo) GetUsers(ctx context.Context, req *users.ReadAllReq) (*users.ReadAllRes, error) {
	rs, err := u.usersClient.ReadAll(ctx, req)
	if err != nil {
		return nil, err
	}

	res, err := rs.Recv()
	if err != nil {
		return nil, err
	}
	return res, err
}

func (u *UserRepo) CreateUser(ctx context.Context, req *users.CreateUserReq) (res *users.CreateUserRes, err error) {
	res, err = u.usersClient.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return
}
