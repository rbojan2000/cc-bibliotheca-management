package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Borrow struct {
	ID         string             `json:"id"`
	Membership string             `json:"membership"`
	City       string             `json:"city"`
	Book       Book               `json:"book"`
	Date       primitive.DateTime `json:"date"`
}
