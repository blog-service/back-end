package responses

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserSignInResponse struct {
	Token string `json:"token"`
}

type UserGetInfoResponse struct {
	Username  string              `json:"username"`
	Email     string              `json:"email"`
	FirstName string              `json:"firstName"`
	LastName  string              `json:"lastName"`
	Phone     string              `json:"phone"`
	RoleId    int                 `json:"role"`
	RoleName  string              `json:"roleName"`
	CreatedAt time.Time           `json:"createdAt"`
	CreatedBy *primitive.ObjectID `json:"createdBy"`
	UpdatedAt time.Time           `json:"updatedAt"`
	UpdatedBy *primitive.ObjectID `json:"updatedBy"`
}
