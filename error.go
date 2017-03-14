package todong

// Error struct
type Error struct {
	Message string `json:"message"`
}

// ErrorResponse struct
type ErrorResponse struct {
	Error *Error `json:"error"`
}

// NewError func
func NewError(error string) *ErrorResponse {
	return &ErrorResponse{&Error{error}}
}
