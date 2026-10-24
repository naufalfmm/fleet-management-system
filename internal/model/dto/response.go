package dto

type ErrorResponse struct {
	Error error `json:"error"`
}

type Response[T any] struct {
	Ok      bool   `json:"ok"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
