package constants

var (
	ErrLoadConfig     = "failed to load config file"
	ErrParseConfig    = "failed to parse env to config struct"
	ErrUnknown        = "unknown error"
	ErrMissingToken   = "missing token"
	ErrWrongToken     = "wrong token"
	ErrInvalidRequest = "invalid request"
)

const (
	ErrCodeUnknown        = 1000
	ErrCodeMissingToken   = 1001
	ErrCodeWrongToken     = 1002
	ErrCodeInvalidRequest = 1003
)
