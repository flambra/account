package internal

import (
	"os"

	"github.com/flambra/account/internal/auth"
	"github.com/flambra/account/internal/auth/twoFactor"
	"github.com/flambra/account/internal/middleware"
	"github.com/flambra/account/internal/profile"
	"github.com/flambra/account/internal/user"
	"github.com/gofiber/fiber/v2"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", middleware.Auth, func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	app.Post("/user", middleware.Auth, user.Create)
	app.Get("/user/:id", middleware.Auth, user.Read)
	app.Put("/user/:id", middleware.Auth, user.Update)
	app.Delete("/user/:id", middleware.Auth, user.Delete)

	app.Get("/profile/:id", middleware.Auth, profile.Read)
	app.Put("/profile/:id", middleware.Auth, profile.Update)
	app.Get("/profile", middleware.Auth, profile.Find)

	app.Post("/auth/signin", middleware.Auth, auth.SignIn)
	
	app.Post("/auth/twofactor/send", middleware.Auth, twoFactor.Send)
	app.Post("/auth/twofactor/validate", middleware.Auth, twoFactor.Validate)
}
