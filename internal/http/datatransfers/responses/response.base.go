package responses

type BaseResponse struct {
	Success   bool        `json:"success"`
	Message   string      `json:"message,omitempty"`
	ErrorCode int         `json:"error_code,omitempty"`
	Data      interface{} `json:"data,omitempty"`
}
