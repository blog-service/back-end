package requests

import (
	"errors"
	"regexp"

	"back-end/pkg/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserGetUserByIdRequest struct {
	Id primitive.ObjectID `query:"id"`
}

type UserSignUpRequest struct {
	Username  string `json:"username" validate:"required,alphanum"`
	Password  string `json:"password" validate:"required"`
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Phone     string `json:"phone" validate:"required,e164"`
}

func (r *UserSignUpRequest) Validate() error {
	if err := validator.New().ValidatePayloads(r); err != nil {
		return err
	}
	re := regexp.MustCompile(`^[A-Za-z\d@$!%*#?&]{6,}$`)
	if !re.MatchString(r.Password) {
		return errors.New("password must contain at least six characters, at least one letter, one number and one special character")
	}
	return nil
}
