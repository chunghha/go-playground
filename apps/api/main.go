package main

import (
	"go-playground/apps/api/pkg/configs"
	"go-playground/apps/api/pkg/middleware"
	"go-playground/apps/api/pkg/routes"
	"go-playground/apps/api/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"

	_ "go-playground/docs" // load API Docs files (Swagger)

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

var logger *zap.Logger

func initLogger() {
	logger, _ = zap.NewProduction()
	defer logger.Sync()
	zap.ReplaceGlobals(logger)
}

// @title go-playground
// @version 0.1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	initLogger()

	config := configs.FiberConfig()
	app := fiber.New(config)

	middleware.FiberMiddleware(app)

	routes.HealthRoute(app)
	routes.PublicRoutes(app)
	routes.SwaggerRoute(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, Flutter 👋!")
	})

	routes.NotFoundRoute(app)

	utils.StartServer(app)
}
