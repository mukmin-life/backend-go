package main

import (

	_ "github.com/mukmin-life/docs"
	"github.com/mukmin-life/db"
	prayer_api "github.com/mukmin-life/modules/prayer_time"
	"context"

	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/swaggo/fiber-swagger"

	texporter "github.com/GoogleCloudPlatform/opentelemetry-operations-go/exporter/trace"
	"go.opentelemetry.io/contrib/detectors/gcp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
)

// @title Mukmin Life API
// @version 1.0

// @contact.name Mukmin Life API Support
// @contact.email hello@mukmin.life

// @license.name Proprietary

// @host localhost:3000
// @BasePath /v2
func main() {

	// Initiate OTEL tracing
	// Ref: https://cloud.google.com/trace/docs/setup/go-ot
	ctx := context.Background()
	projectID := os.Getenv("GOOGLE_CLOUD_PROJECT")
	exporter, err := texporter.New(texporter.WithProjectID(projectID))
	if err == nil {
		// Identify your application using resource detection
		res, err := resource.New(ctx,
				resource.WithDetectors(gcp.NewDetector()),
				resource.WithTelemetrySDK(),
				resource.WithAttributes(
						semconv.ServiceNameKey.String("mukmin-api"),
				),
		)
		if err != nil {
				log.Fatalf("resource.New: %v", err)
		}

		tp := sdktrace.NewTracerProvider(
				sdktrace.WithBatcher(exporter),
				sdktrace.WithResource(res),
		)
		defer tp.ForceFlush(ctx)
		otel.SetTracerProvider(tp)
	}

    err = db.Connect()
    if err != nil {
        log.Fatal("failed to connect to DB")
    }
    defer db.Connection.Close()

	app := fiber.New()

	app.Get("/docs/*", swagger.WrapHandler)
	api := app.Group("/")
	v2 := api.Group("/v2")

	prayer_time := v2.Group("/prayer_time")
	prayer_time.Get("/:zone/:date", prayer_api.GetPrayerTime)

    log.Fatal(app.Listen(":3000"))
}
