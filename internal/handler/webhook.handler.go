package handler

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type webHookHandler struct {
}

func (wh *webHookHandler) ReceiveInvoice(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	return c.SendStatus(http.StatusOK)
}

func NewWebHookHandler() *webHookHandler {
	return &webHookHandler{}
}
