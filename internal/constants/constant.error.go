package constants

var (
	ErrUnknown        = "unknown error"
	ErrMissingToken   = "missing token"
	ErrWrongToken     = "wrong token"
	ErrInvalidRequest = "invalid request"
)

const (
	ErrCodeNoErr              = 0
	ErrCodeUnknown            = 1000
	ErrCodeMissingToken       = 1001
	ErrCodeWrongToken         = 1002
	ErrCodeParseRequestFailed = 1003
	ErrCodeInvalidRequest     = 1004

	ErrCodeHashPassFailed = 1005
	ErrCodeDuplicateData  = 1006
)
