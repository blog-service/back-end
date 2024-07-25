package requests

import (
	"back-end/pkg/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGetUserByIdRequest struct {
	Id primitive.ObjectID `query:"id"`
}

type UserSignUpRequest struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone" validate:"required"`
}

func (r *UserSignUpRequest) Validate() error {
	if err := validator.New().ValidatePayloads(r); err != nil {
		return err
	}
	return nil
}
