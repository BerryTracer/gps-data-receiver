package api

import "github.com/gofiber/fiber/v2"

type FiberContext struct {
	Ctx *fiber.Ctx
}

func (c *FiberContext) Bind(i interface{}) error {
	return c.Ctx.BodyParser(i)
}

func (c *FiberContext) JSON(statusCode int, i interface{}) error {
	return c.Ctx.Status(statusCode).JSON(i)
}
