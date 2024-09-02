package user

import (
	"github.com/MauroPerna/go-fiber-backend-test/services/database"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type UserModule struct {
	UserService *UserService
}

func NewUserModule(databaseModule *database.DatabaseModule) *UserModule {
	userService := NewUserService(databaseModule)
	return &UserModule{
		UserService: userService,
	}
}

func (m *UserModule) SetupRoutes(app *fiber.App) {
	app.Get("/users", m.UserService.GetAllUsers)
	app.Get("/users/:id", m.UserService.GetUser)
}

var Module = fx.Options(
	fx.Provide(NewUserModule),
)
