package handler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

func UploadProducrtImageHandler(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "image data not found",
		})
	}

	// validasi gambar (ekstensi dan content type)

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".webp": true,
	}
	if !allowedExts[ext] {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "image extension is not allowed (jpg, jpeg, png, and webp)",
		})
	}

	// validasi content type
	contentType := file.Header.Get("Content-Type")
	allowedContentType := map[string]bool{
		"image/jpeg": true,
		"image/png":  true,
		"image/webp": true,
	}
	if !allowedContentType[contentType] {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{
			"success": false,
			"message": "this content type is not allowed (image/jpeg, image/png, image/webp)",
		})
	}

	timestamp := time.Now().UnixNano()
	fileName := fmt.Sprintf("/product_%d%s", timestamp, filepath.Ext(file.Filename))
	uploadPath := "./storage/product/" + fileName
	if err := c.SaveFile(file, uploadPath); err != nil {
		fmt.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"success": false,
			"message": "internal server error",
		})
	}

	return c.JSON(fiber.Map{
		"success":   true,
		"message":   "Upload success",
		"file_name": fileName,
	})
}
