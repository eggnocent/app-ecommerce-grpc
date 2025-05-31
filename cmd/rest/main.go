package main

import (
	"context"
	"github/eggnocent/app-grpc-eccomerce/internal/handler"
	"github/eggnocent/app-grpc-eccomerce/internal/repository"
	"github/eggnocent/app-grpc-eccomerce/internal/service"
	"github/eggnocent/app-grpc-eccomerce/pkg/database"
	"log"
	"mime"
	"net/http"
	"os"
	"path"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func handleGetFileName(c *fiber.Ctx) error {
	fileNameParam := c.Params("filename")
	filePath := path.Join("storage", "product", fileNameParam)
	if _, err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			return c.Status(http.StatusNotFound).SendString("Not Found")
		}

		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal server errror")
	}

	// buka file
	file, err := os.Open(filePath)
	if err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).SendString("Internal server errror")
	}

	ext := path.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	c.Set("Content-Type", mimeType)
	// kirim file sbg response
	return c.SendStream(file)
}

func main() {
	godotenv.Load()
	ctx := context.Background()
	app := fiber.New()

	db := database.ConnectDB(ctx, os.Getenv("DB_URI"))

	orderRepository := repository.NewOrderRepository(db)
	webHookService := service.NewWebHookService(orderRepository)
	webHookHandler := handler.NewWebHookHandler(webHookService)

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/storage/product/:filename", handleGetFileName)

	app.Post("/product/upload", handler.UploadProducrtImageHandler)
	app.Post("/webhook/xendit/invoice", webHookHandler.ReceiveInvoice)

	app.Listen(":3000")
}
