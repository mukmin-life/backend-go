package main

import (
	"github.com/mukmin-life/db"
	_ "github.com/mukmin-life/docs"
	prayer_api "github.com/mukmin-life/modules/prayer_time"

	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	swagger "github.com/swaggo/fiber-swagger"
)

// @title Mukmin Life API
// @version 1.0

// @contact.name Mukmin Life API Support
// @contact.email hello@mukmin.life

// @license.name Proprietary

// @host localhost:3000
// @BasePath /v2
func main() {
    err := db.Connect()
    if err != nil {
        log.Fatal("failed to connect to DB")
    }
    defer db.Connection.Close()

	app := fiber.New()

// Initialize default config
app.Use(cors.New())

// Or extend your config for customization
app.Use(cors.New(cors.Config{
    AllowOrigins: "*",
}))

	app.Get("/docs/*", swagger.WrapHandler)
	api := app.Group("/")
	v2 := api.Group("/v2")

	prayer_time := v2.Group("/prayer_time")
	prayer_time.Get("/:zone/:date", prayer_api.GetPrayerTime)

	log.Fatal(app.Listen(":3000"))
}
