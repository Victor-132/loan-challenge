package httpserver

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type FiberAdpater struct {
	app *fiber.App
}

func NewFiberAdapter() HttpServer[*fiber.Ctx] {
	return &FiberAdpater{app: fiber.New()}
}

// Listen implements HttpServer.
func (f *FiberAdpater) Listen(port int) {
	f.app.Listen(fmt.Sprintf(":%d", port))
}

// On implements HttpServer.
func (f *FiberAdpater) On(method, path string, callback func(*fiber.Ctx) error) {
	f.app.Add(method, path, callback)
}

// Test implements HttpServer.
func (f *FiberAdpater) Test(req *http.Request) (*http.Response, error) {
	return f.app.Test(req)
}
