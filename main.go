package main

import (
	"log"

	"github.com/aiteung/musik"
	"github.com/ryaasishlah/websocket-herokuyas/module"
	"github.com/ryaasishlah/websocket-herokuyas/url"

	"github.com/gofiber/fiber/v2"
)

func main() {
	go module.RunHub()

	site := fiber.New()
	url.Web(site)
	log.Fatal(site.Listen(musik.Dangdut()))
}
