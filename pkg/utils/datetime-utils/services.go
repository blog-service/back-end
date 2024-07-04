package datetime_utils

import "time"

func (s *service) Format(t *time.Time) string {
	if t == nil {
		return ""
	}
	if t.IsZero() {
		return ""
	}
	return t.Local().Format("2006-01-02 15:04:05 Z07")
}
