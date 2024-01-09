package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rbojan2000/central-library/model"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r repository) GetUser(ctx context.Context, id string) (model.User, error) {
	var out user
	err := r.db.
		Collection("users").
		FindOne(ctx, bson.M{"id": id}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.User{}, ErrUserNotFound
		}
		return model.User{}, err
	}
	return toModel(out), nil
}

func (r repository) CreateUser(ctx context.Context, user model.User) (model.User, error) {
	out, err := r.db.
		Collection("users").
		InsertOne(ctx, fromModel(user))
	if err != nil {
		return model.User{}, err
	}
	user.ID = out.InsertedID.(primitive.ObjectID).String()
	return user, nil
}

func (r repository) UpdateUser(ctx context.Context, user model.User) (model.User, error) {
	in := bson.M{}
	in["numOfRentedBooks"] = user.NumOfRentedBooks
	out, err := r.db.
		Collection("users").
		UpdateOne(ctx, bson.M{"id": user.ID}, bson.M{"$set": in})
	if err != nil {
		return model.User{}, err
	}
	if out.MatchedCount == 0 {
		return model.User{}, ErrUserNotFound
	}
	return user, nil
}

func (r repository) DeleteUser(ctx context.Context, id string) error {
	out, err := r.db.
		Collection("users").
		DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrUserNotFound
	}
	return nil
}

type user struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Surname          string `json:"surname"`
	NumOfRentedBooks int    `json:"numOfRentedBooks"`
}

func fromModel(in model.User) user {
	return user{
		Name:             in.Name,
		Surname:          in.Surname,
		ID:               in.ID,
		NumOfRentedBooks: in.NumOfRentedBooks,
	}
}

func toModel(in user) model.User {
	return model.User{
		ID:               in.ID,
		Name:             in.Name,
		Surname:          in.Surname,
		NumOfRentedBooks: in.NumOfRentedBooks,
	}
}
