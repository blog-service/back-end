package validator

type service struct {
}

type Service interface {
	ValidatePayloads(payload interface{}) (err error)
}

func New() Service {
	return &service{}
}

var mapHelper = map[string]string{
	"required":  "is a required field",
	"email":     "is not a valid email address",
	"lowercase": "must contain at least one lowercase letter",
	"uppercase": "must contain at least one uppercase letter",
	"numeric":   "must contain at least one digit",
}

var needParam = []string{"min", "max", "containsany"}
