package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	Id        primitive.ObjectID  `bson:"_id,omitempty"`
	Username  string              `bson:"username"`
	Email     string              `bson:"email"`
	FirstName string              `bson:"first_name"`
	LastName  string              `bson:"last_name"`
	Password  string              `bson:"password"`
	Phone     string              `bson:"phone"`
	Status    int                 `bson:"status"`
	RoleId    int                 `bson:"role"`
	CreatedAt time.Time           `bson:"created_at"`
	CreatedBy *primitive.ObjectID `bson:"created_by"`
	UpdatedAt time.Time           `bson:"updated_at"`
	UpdatedBy *primitive.ObjectID `bson:"updated_by"`
	DeletedAt *time.Time          `bson:"deleted_at"`
	DeletedBy *primitive.ObjectID `bson:"deleted_by"`
}
