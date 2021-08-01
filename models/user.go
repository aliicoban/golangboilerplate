package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
  )

type Users struct {
	ID    primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Email string `json:"email"`
	Password string `json:"password"`
	Role string `json:"role"`
}