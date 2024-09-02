package config

import (
	"log"

	"github.com/MauroPerna/go-fiber-backend-test/services/secret_manager"
	"go.uber.org/fx"
)

type ConfigModule struct {
	ConfigService *ConfigService
}

func NewConfigModule(s *secret_manager.SecretManagerService) *ConfigModule {
	configService, err := NewConfigService(s)
	if err != nil {
		log.Fatalf("Failed to initialize ConfigService: %v", err)
	}

	return &ConfigModule{
		ConfigService: configService,
	}
}

// Modifica el módulo para proveer ConfigService también
var Module = fx.Options(
	fx.Provide(
		NewConfigModule,
		func(m *ConfigModule) *ConfigService { return m.ConfigService }, // Proveer ConfigService
	),
)
