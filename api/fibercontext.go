package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
)

// FiberContext is a struct that wraps the fiber.Ctx context.
// It provides methods for binding request data and sending JSON responses.
type FiberContext struct {
	Ctx *fiber.Ctx // Ctx is the context from the fiber framework.
}

// NewFiberContext is a constructor for the FiberContext struct.
func NewFiberContext(ctx *fiber.Ctx) *FiberContext {
	return &FiberContext{Ctx: ctx}
}

// Bind is a method on the FiberContext struct.
// It uses the BodyParser method of the fiber context to parse the request body into the provided struct.
// It returns an error if the parsing fails.
func (c *FiberContext) Bind(i interface{}) error {
	return c.Ctx.BodyParser(i)
}

// JSON is a method on the FiberContext struct.
// It sets the HTTP status code and sends a JSON response with the provided payload.
// It returns an error if the operation fails.
func (c *FiberContext) JSON(statusCode int, i interface{}) error {
	return c.Ctx.Status(statusCode).JSON(i)
}

// SendStatus sends the specified HTTP status code as a response.
// It uses the underlying Fiber context to send the status code.
// The statusCode parameter specifies the HTTP status code to be sent.
// Returns an error if there was an issue sending the status code.
func (c *FiberContext) SendStatus(statusCode int) error {
	return c.Ctx.SendStatus(statusCode)
}

// Context returns the context.Context from the underlying framework context.
func (c *FiberContext) Context() context.Context {
	return c.Ctx.Context()
}
