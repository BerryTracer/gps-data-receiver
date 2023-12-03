package api

import "github.com/gofiber/fiber/v2"

type HttpRouter interface {
	Post(string, func(HttpContext) error)
	Get(string, func(HttpContext) error)
}

type FiberRouter struct {
	App *fiber.App
}

func NewFiberRouter(app *fiber.App) *FiberRouter {
	return &FiberRouter{App: app}
}

func (f *FiberRouter) Post(path string, handler func(HttpContext) error) {
	f.App.Post(path, func(c *fiber.Ctx) error {
		return handler(NewFiberContext(c))
	})
}

func (f *FiberRouter) Get(path string, handler func(HttpContext) error) {
	f.App.Get(path, func(c *fiber.Ctx) error {
		return handler(NewFiberContext(c))
	})
}

// Ensure that FiberRouter implements HttpRouter
var _ HttpRouter = (*FiberRouter)(nil)
