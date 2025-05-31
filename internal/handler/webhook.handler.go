package handler

import (
	"github/eggnocent/app-grpc-eccomerce/internal/dto"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type webHookHandler struct {
	webHookService service.IwebHookService
}

func (wh *webHookHandler) ReceiveInvoice(c *fiber.Ctx) error {
	var request dto.XenditInvoiceRequest
	body := c.Body()
	log.Println("Webhook body:", string(body))

	err := c.BodyParser(&request)
	if err != nil {
		log.Println("Parse error:", err)
		return c.SendStatus(http.StatusBadRequest)
	}

	err = wh.webHookService.ReceiveInvoice(c.UserContext(), &request)
	if err != nil {
		log.Println(err)
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}

func NewWebHookHandler(webHookService service.IwebHookService) *webHookHandler {
	return &webHookHandler{
		webHookService: webHookService,
	}
}
