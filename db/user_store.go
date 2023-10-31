package db

import (
	"context"
	"github.com/samarthasthan/go-hotel/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const DBColl = "users"

type UserStore interface {
	GetUserByID(context.Context, string) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
}

type MongoUserStore struct {
	client *mongo.Client
	coll   *mongo.Collection
}

func NewMongoUserStore(client *mongo.Client) *MongoUserStore {
	return &MongoUserStore{
		coll: client.Database(DBName).Collection(DBColl),
	}
}

func (s *MongoUserStore) InsertUser(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, &user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID)
	return user, nil
}

func (s *MongoUserStore) GetUserByID(ctx context.Context, id string) (*types.User, error) {
	var user types.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = s.coll.FindOne(ctx, bson.M{"_id": oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *MongoUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {

	results, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return []*types.User{}, err
	}
	var users []*types.User
	err = results.All(ctx, &users)
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return []*types.User{}, nil
	}
	return users, nil
}
