package database

import (
	"github.com/MauroPerna/go-fiber-backend-test/services/config"
	"go.uber.org/fx"
)

type DatabaseModule struct {
	DatabaseService *DatabaseService
}

func NewDatabaseModule(lc fx.Lifecycle, c *config.ConfigService) *DatabaseModule {
	return &DatabaseModule{
		DatabaseService: NewDatabaseService(lc, c),
	}
}

var Module = fx.Options(
	fx.Provide(NewDatabaseModule),
)
