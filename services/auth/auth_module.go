package auth

import (
	"github.com/MauroPerna/go-fiber-backend-test/services/config"
	"go.uber.org/fx"
)

type AuthModule struct {
	AuthService *AuthService
}

func NewAuthModule(c *config.ConfigService) *AuthModule {
	return &AuthModule{
		AuthService: NewAuthService(c),
	}
}

var Module = fx.Options(
	fx.Provide(NewAuthModule),
)
