package string_utils

type service struct{}

type Service interface {
	IsArrayContains(arr []string, str string) bool
}

func New() Service {
	return &service{}
}
