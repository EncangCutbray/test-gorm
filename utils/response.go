package utils

type SuccessResponse struct {
	Code    int         `json:"code"`
	Status  string      `json:"status"` // misal: "success"
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code    int               `json:"code"`
	Status  string            `json:"status"` // misal: "fail" atau "error"
	Message string            `json:"message"`
	Errors  map[string]string `json:"errors,omitempty"` // detail error opsional
}
