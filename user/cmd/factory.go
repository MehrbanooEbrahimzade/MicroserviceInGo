package cmd

import (
	"context"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/config"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/handler"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/mongodb"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/repo"
	"github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/service"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
)

type Factory struct {
	Config      *config.Config
	Db          *mongodb.Mongodb
	GrpcServer  *grpc.Server
	GrpcHandler *handler.UserHandler
}

func NewFactory(ctx context.Context) (*Factory, error) {
	config, err := config.LoadConfig()
	if err != nil {
		log.Println(err)
	}

	db := mongodb.Initialdb(config.MongodbUrl)
	userColl := addCollection(db.Client, ctx)

	//Initializing layers
	repo, err := repo.NewUserRepository(userColl)
	if err != nil {
		log.Println(err)
	}
	service, err := service.NewService(repo)
	if err != nil {
		log.Println(err)
	}
	handler, err := handler.NewHandler(service)
	if err != nil {
		log.Println(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	return &Factory{
		Config:      config,
		Db:          db,
		GrpcServer:  grpcServer,
		GrpcHandler: handler,
	}, nil
}
func addCollection(client *mongo.Client, ctx context.Context) *mongo.Collection {
	userColl := client.Database("mydb").Collection("user")
	mod := mongo.IndexModel{Keys: bson.M{"mobileNo": 1}, Options: options.Index().SetUnique(true)}
	userColl.Indexes().CreateOne(ctx, mod)
	return userColl
}
