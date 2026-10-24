package dto

type ErrorResponse struct {
	Error error `json:"error"`
}

func (e ErrorResponse) Args() []any {
	return []any{"error", e.Error.Error()}
}

type Response[T any] struct {
	Ok      bool   `json:"ok"`
	Code    int    `json:"code,omitempty"`
	Message string `json:"message"`
	Data    T      `json:"data"`
}
