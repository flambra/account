package internal

import (
	"os"

	"github.com/flambra/account/internal/auth"
	"github.com/flambra/account/internal/auth/twoFactor"
	"github.com/flambra/account/internal/profile"
	"github.com/flambra/account/internal/user"
	"github.com/flambra/helpers/hMiddleware"
	"github.com/flambra/helpers/hToken"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func InitializeRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"project":     os.Getenv("PROJECT"),
			"environment": os.Getenv("ENV"),
			"version":     os.Getenv("BUILD_VERSION"),
		})
	})

	// Docs
	app.Get("/swagger/*", hMiddleware.BasicAuth(), swagger.HandlerDefault)

	// Auth
	app.Post("/auth/signin", auth.SignIn)
	app.Post("/auth/refreshtoken", auth.RefreshToken)

	app.Post("/user", hToken.Middleware, user.Create)
	app.Get("/user/:id", hToken.Middleware, user.Read)
	app.Put("/user/:id", hToken.Middleware, user.Update)
	app.Delete("/user/:id", hToken.Middleware, user.Delete)
	app.Get("/user", hToken.Middleware, user.Page)

	app.Get("/profile/:id", hToken.Middleware, profile.Read)
	app.Put("/profile/:id", hToken.Middleware, profile.Update)
	app.Get("/profile", hToken.Middleware, profile.Find)

	app.Post("/auth/twofactor/send", hToken.Middleware, twoFactor.Send)
	app.Post("/auth/twofactor/validate", hToken.Middleware, twoFactor.Validate)
}
