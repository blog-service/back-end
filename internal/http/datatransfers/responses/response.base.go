package responses

type ErrorResponse struct {
	Success   bool   `json:"success"`
	Message   string `json:"message,omitempty"`
	ErrorCode int    `json:"error_code,omitempty"`
}

type SuccessResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}
