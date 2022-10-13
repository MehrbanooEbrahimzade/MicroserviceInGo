package cmd

import (
	"context"
	"fmt"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/configs"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/handler"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/models"
	userRepo "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/repository/user"
	userSvc "github.com/MehrbanooEbrahimzade/MicroserviceInGo/gateway/pkg/service/user"
	"net/http"
)

type factory struct {
	Config  models.Config
	Port    string
	Handler http.Handler
}

func NewFactory() *factory {
	configs := configs.NewConfig()
	port := fmt.Sprintf(":%d", configs.Port)

	userClient := userRepo.NewUserRepoClient(configs.UsersPort)
	ur := userRepo.NewUsersRepo(*userClient)
	us := userSvc.NewUserSvc(context.Background(), ur)

	h := handler.NewHandler(us)
	router := handler.NewHandlerRoute(configs.Cors, h)
	return &factory{
		Config:  *configs,
		Port:    port,
		Handler: router,
	}
}
