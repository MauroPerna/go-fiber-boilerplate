package app

import (
	middleware "github.com/MauroPerna/go-fiber-backend-test/common/middlewares"
	"github.com/MauroPerna/go-fiber-backend-test/domain/user"
	"github.com/MauroPerna/go-fiber-backend-test/services/auth"
	"github.com/MauroPerna/go-fiber-backend-test/services/config"
	"github.com/MauroPerna/go-fiber-backend-test/services/database"
	"github.com/MauroPerna/go-fiber-backend-test/services/secret_manager"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type AppModule struct {
	FiberApp            *fiber.App
	ConfigModule        *config.ConfigModule
	SecretManagerModule *secret_manager.SecretManagerModule
	DatabaseModule      *database.DatabaseModule
	AuthModule          *auth.AuthModule
	UserModule          *user.UserModule
}

func NewAppModule(
	lc fx.Lifecycle,
	userModule *user.UserModule,
	configModule *config.ConfigModule,
	authModule *auth.AuthModule,
	secretManagerModule *secret_manager.SecretManagerModule,
	databaseModule *database.DatabaseModule,
) *AppModule {
	app := fiber.New()

	// Registrar el middleware globalmente
	app.Use(middleware.AuthMiddleware(authModule.AuthService))
	app.Use(middleware.LoggerMiddleware())

	// Configurar las rutas
	userModule.SetupRoutes(app)

	return &AppModule{
		FiberApp:            app,
		AuthModule:          authModule,
		UserModule:          userModule,
		ConfigModule:        configModule,
		SecretManagerModule: secretManagerModule,
		DatabaseModule:      databaseModule,
	}
}

var Module = fx.Options(
	fx.Provide(NewAppModule),
)
