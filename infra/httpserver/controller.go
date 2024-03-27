package httpserver

import (
	"net/http"

	"github.com/Victor-132/loan-challenge.git/application/usecase"
	"github.com/gofiber/fiber/v2"
)

func NewController(hs HttpServer[*fiber.Ctx], uc usecase.UseCase[usecase.Input, usecase.Output]) {
	hs.On(http.MethodPost, "/customer-loans", func(c *fiber.Ctx) error {
		var input usecase.Input

		if err := c.BodyParser(&input); err != nil {
			return err
		}

		out := uc.Execute(input)

		return c.JSON(out)
	})
}
