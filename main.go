package main

import (
	"context"
	"log"

	app "github.com/MauroPerna/go-fiber-backend-test/domain"
	"github.com/MauroPerna/go-fiber-backend-test/domain/user"
	"github.com/MauroPerna/go-fiber-backend-test/services/auth"
	"github.com/MauroPerna/go-fiber-backend-test/services/config"
	"github.com/MauroPerna/go-fiber-backend-test/services/database"
	"github.com/MauroPerna/go-fiber-backend-test/services/secret_manager"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		auth.Module,
		user.Module,
		config.Module,
		secret_manager.Module,
		database.Module,
		app.Module,
		fx.Invoke(
			startServer, // Invocar la función para iniciar el servidor
		),
	)

	app.Run()
}

func startServer(appModule *app.AppModule, lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			port := appModule.ConfigModule.ConfigService.Get("PORT", "8080")
			log.Printf("server is running on port :%s", port)

			go func() {
				if err := appModule.FiberApp.Listen(":3000"); err != nil {
					log.Fatalf("Failed to start server: %v", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			if err := appModule.FiberApp.Shutdown(); err != nil {
				log.Fatalf("Failed to stop server: %v", err)
			}
			return nil
		},
	})
}
