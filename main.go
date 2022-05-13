package main

import (
	"encoding/json"
	"fmt"
	"github.com/darkCavalier11/downloader_backend/models"
	"log"
	"os/exec"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Hello World!")
	})

	app.Get("/get_file_meta", func(c *fiber.Ctx) error {
		query := struct {
			Url string `json:"url"`
		}{}
		fmt.Println(c.QueryParser(&query))
		err := c.QueryParser(&query)
		if err != nil {
			return fiber.ErrBadRequest
		}
		log.Println(query)
		cmd := exec.Command("yt-dlp", "--dump-json", "--skip-download", query.Url)
		stdout, err := cmd.Output()

		if err != nil {
			return fiber.ErrServiceUnavailable
		}

		var fileMeta models.FileMeta
		err = json.Unmarshal([]byte(stdout), &fileMeta)
		if err != nil {
			log.Println(err)
			return fiber.ErrBadRequest
		}
		_ = stdout
		return c.JSON(fileMeta)
	})

	log.Fatalln(app.Listen(":8000"))
}
