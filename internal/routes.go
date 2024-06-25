package internal

import (
	"os"

	"github.com/flambra/account/internal/auth"
	"github.com/flambra/account/internal/auth/twoFactor"
	"github.com/flambra/account/internal/profile"
	"github.com/flambra/account/internal/user"
	"github.com/flambra/helpers/hMiddleware"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", middleware.Auth, func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	// Docs
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Post("/user", hMiddleware.Auth, user.Create)
	app.Get("/user/:id", hMiddleware.Auth, user.Read)
	app.Put("/user/:id", hMiddleware.Auth, user.Update)
	app.Delete("/user/:id", hMiddleware.Auth, user.Delete)
	app.Get("/users/page", hMiddleware.Auth, user.Page)

	app.Get("/profile/:id", hMiddleware.Auth, profile.Read)
	app.Put("/profile/:id", hMiddleware.Auth, profile.Update)
	app.Get("/profile", hMiddleware.Auth, profile.Find)

	app.Post("/auth/signin", hMiddleware.Auth, auth.SignIn)

	app.Post("/auth/twofactor/send", hMiddleware.Auth, twoFactor.Send)
	app.Post("/auth/twofactor/validate", hMiddleware.Auth, twoFactor.Validate)
}
