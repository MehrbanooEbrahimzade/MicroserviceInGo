package repo

import (
	"context"
	"fmt"
	pb "github.com/MehrbanooEbrahimzade/MicroserviceInGo/user/pkg/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

type UserItem struct {
	id        primitive.ObjectID `bson:"_id,omitempty"`
	userName  string             `bson:"userName"`
	email     string             `bson:"email"`
	mobileNo  string             `bson:"mobileNo"`
	birthDate string             `bson:"birthDate"`
	password  string             `bson:"password"`
}
type userRepoImpl struct {
	db *mongo.Collection
}

func NewUserRepository(db *mongo.Collection) (*userRepoImpl, error) {
	return &userRepoImpl{
		db: db,
	}, nil
}

func (repo *userRepoImpl) CreateRepo(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	result, err := repo.db.InsertOne(ctx, req)
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return &pb.CreateUserRes{
		ID: oid.Hex(),
	}, nil
}

func (repo *userRepoImpl) ReadAllRepo(req *pb.ReadAllReq, stream pb.UserService_ReadAllServer) error {
	data := &UserItem{}
	cursor, err := repo.db.Find(context.Background(), bson.M{})
	if err != nil {
		log.Fatalf("error in read all repo")
		return err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		err := cursor.Decode(data)
		if err != nil {
			return status.Errorf(codes.Unavailable, fmt.Sprintf("Could not decode data: %v", err))
		}
		stream.Send(&pb.ReadAllRes{
			User: &pb.User{
				ID:        data.id.Hex(),
				UserName:  data.userName,
				MobileNo:  data.mobileNo,
				Email:     data.email,
				BirthDate: data.birthDate,
				Password:  data.password,
			},
		})
	}
	if err := cursor.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unkown cursor error: %v", err))
	}
	return nil
}
