package entities

type GetHealthUserRequest struct {
	Service string `json:"service"`
}

type GetHealthUserResponse struct {
	Status string `json:"status"`
}
