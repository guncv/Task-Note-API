package entities

type ErrorResponse struct {
	Code    int    `json:"code" example:"1001"`
	Message string `json:"message" example:"Validation failed"`
}
