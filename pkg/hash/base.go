package hash

type service struct {
}

type Service interface {
	HashPassword(password string) (hashPass string, err error)
	CheckPasswordHash(password, hash string) bool
}

func New() Service {
	return &service{}
}
