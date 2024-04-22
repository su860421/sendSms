package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type RequestData struct {
	MobileNumber string `json:"mobile_number" validate:"required,numeri"`
	Title        string `json:"title" validate:"required,string"`
	Message      string `json:"message" validate:"required,text"`
}

func main() {
	app := fiber.New()

	// 設定路由
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/send-message", func(c *fiber.Ctx) error {
		requestData := new(RequestData)

		if err := c.BodyParser(requestData); err != nil {
			return err
		}

		validate := validator.New()
		if err := validate.Struct(requestData); err != nil {
			return err
		}

		return c.SendString("Hello, World!")
	})

	// 啟動伺服器
	err := app.Listen(":8080")
	if err != nil {
		panic(err)
	}
}
