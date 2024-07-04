package datetime_utils

import "time"

type service struct{}

type Service interface {
	Format(t *time.Time) string
}

func New() Service {
	return &service{}
}
