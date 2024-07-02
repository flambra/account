package main

import (
	"fmt"
	"log"
	"os"

	_ "github.com/flambra/account/docs"
	"github.com/flambra/account/internal"
	"github.com/flambra/account/internal/domain"
	"github.com/flambra/helpers/hDb"
	"github.com/gofiber/contrib/fibernewrelic"
	"github.com/gofiber/contrib/fiberzerolog"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		err = nil
		godotenv.Load("../.env")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = hDb.New()
	if err != nil {
		log.Fatal(err)
	}

	err = hDb.Migrate(
		&domain.User{},
		&domain.Profile{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate the database: %v", err)
		return
	}
}

func main() {
	app := fiber.New()

	fiber.SetParserDecoder(fiber.ParserConfig{
		IgnoreUnknownKeys: true,
		ZeroEmpty:         true,
	})

	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	cfg := fibernewrelic.Config{
		License: os.Getenv("NEWRELIC_LICENSE"),
		AppName: os.Getenv("NEWRELIC_APP_NAME"),
		Enabled: os.Getenv("NEWRELIC_ENABLED") == "true",
	}

	app.Use(fiberzerolog.New(fiberzerolog.Config{
		Logger: &logger,
	}))

	app.Use(fibernewrelic.New(cfg))
	if os.Getenv("NEWRELIC_ENABLED") == "true" {
		logger.Info().Msg("NewRelic enabled")
	} else {
		logger.Info().Msg("NewRelic disabled")
	}

	internal.InitializeRoutes(app)

	port := os.Getenv("SERVER_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	/* Start Server */
	err := app.Listen(fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatal(err)
	}
}
