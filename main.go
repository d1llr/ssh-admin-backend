package main

import (
	"os"

	controllers2 "github.com/create-go-app/fiber-go-template/app/controllers"
	"github.com/create-go-app/fiber-go-template/app/repositories"
	"github.com/create-go-app/fiber-go-template/app/services"
	"github.com/create-go-app/fiber-go-template/app/services/impl"
	_ "github.com/create-go-app/fiber-go-template/docs" // load API Docs files (Swagger)
	"github.com/create-go-app/fiber-go-template/pkg/configs"
	"github.com/create-go-app/fiber-go-template/pkg/middleware"
	"github.com/create-go-app/fiber-go-template/pkg/routes"
	"github.com/create-go-app/fiber-go-template/pkg/utils"
	"github.com/gofiber/fiber/v2"

	_ "github.com/joho/godotenv/autoload" // load .env file automatically
)

// @title API
// @version 1.0
// @description This is an auto-generated API Docs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your@mail.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @BasePath /api
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	config := configs.FiberConfig()
	app := fiber.New(config)

	repositories := repositories.NewRepositories()
	services := services.Services{
		Auth:  impl.NewAuthService(repositories),
		User:  impl.NewUserService(repositories),
		Ssh:   impl.NewSshService(repositories),
		Token: impl.NewTokenService(repositories),
	}
	controllers := controllers2.NewControllers(services)

	middleware.FiberMiddleware(app) // Register Fiber's middleware for app.

	routes.SwaggerRoute(app)
	routes.PublicRoutes(app, controllers)
	routes.PrivateRoutes(app, controllers)
	routes.NotFoundRoute(app)

	if os.Getenv("STAGE_STATUS") == "dev" {
		utils.StartServer(app)
	} else {
		utils.StartServerWithGracefulShutdown(app)
	}
}
