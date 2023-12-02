package api

import "context"

// HttpContext represents an interface for handling HTTP requests.
// It provides methods for binding request data and sending JSON responses.
type HttpContext interface {
	// Bind binds the passed struct instance to the values in the context's request.
	// It returns an error if any.
	Bind(interface{}) error

	// JSON sends a JSON response with status code and payload.
	// It returns an error if any.
	JSON(int, interface{}) error

	// SendStatus sends a response with the given status code.
	SendStatus(int) error

	// Context returns the context.Context from the underlying framework context.
	Context() context.Context
}
