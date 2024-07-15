package constants

import "errors"

var (
	ErrLoadConfig     = errors.New("failed to load config file")
	ErrParseConfig    = errors.New("failed to parse env to config struct")
	ErrUnknown        = errors.New("unknown error")
	ErrInvalidRequest = errors.New("invalid request")
)

const (
	ErrCodeUnknown = iota + 1000
	ErrCodeInvalidRequest
)
