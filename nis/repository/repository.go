package repository

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/rbojan2000/nis/model"
)

var (
	ErrBorrowNotFound = errors.New("borrow not found")
)

type repository struct {
	db *mongo.Database
}

func NewRepository(db *mongo.Database) Repository {
	return &repository{db: db}
}

func (r repository) GetBorrow(ctx context.Context, id string) (model.Borrow, error) {
	var out borrow
	err := r.db.
		Collection("borrows").
		FindOne(ctx, bson.M{"id": id}).
		Decode(&out)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return model.Borrow{}, ErrBorrowNotFound
		}
		return model.Borrow{}, err
	}
	return toModel(out), nil
}

func (r repository) CreateBorrow(ctx context.Context, borrow model.Borrow) (model.Borrow, error) {
	out, err := r.db.
		Collection("borrows").
		InsertOne(ctx, fromModel(borrow))
	if err != nil {
		return model.Borrow{}, err
	}
	borrow.ID = out.InsertedID.(primitive.ObjectID).String()
	return borrow, nil
}

func (r repository) DeleteBorrow(ctx context.Context, id string) error {
	out, err := r.db.
		Collection("borrows").
		DeleteOne(ctx, bson.M{"id": id})
	if err != nil {
		return err
	}
	if out.DeletedCount == 0 {
		return ErrBorrowNotFound
	}
	return nil
}

type borrow struct {
	ID         string             `json:"id"`
	City       string             `json:"city"`
	Membership string             `json:"membership"`
	Book       model.Book         `json:"book"`
	Date       primitive.DateTime `json:"date"`
}

func fromModel(in model.Borrow) borrow {
	return borrow{
		ID:         in.ID,
		City:       in.City,
		Membership: in.Membership,
		Date:       in.Date,
		Book:       in.Book,
	}
}

func toModel(in borrow) model.Borrow {
	return model.Borrow{
		Membership: in.Membership,
		Book:       in.Book,
		Date:       in.Date,
		City:       in.City,
	}
}
