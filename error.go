package main

// Err struct
type Err struct {
	Message string `json:"message"`
}

// Error struct
type Error struct {
	Error Err `json:"error"`
}

// ErrValidate struct
type ErrValidate struct {
	Message string      `json:"message"`
	Errors  interface{} `json:"errors"`
}

// ErrorValidate struct
type ErrorValidate struct {
	Error ErrValidate `json:"error"`
}

// NewError func
func NewError(err string) *Error {
	return &Error{Error: Err{err}}
}

// NewErrorValidate func
func NewErrorValidate(err error) *ErrorValidate {
	return &ErrorValidate{Error: ErrValidate{"Validation Error", err}}
}
