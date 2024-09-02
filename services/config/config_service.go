package config

import (
	"log"
	"os"

	"github.com/MauroPerna/go-fiber-backend-test/services/secret_manager"
)

type ConfigService struct {
	Secrets map[string]string
}

func NewConfigService(secretManager *secret_manager.SecretManagerService) (*ConfigService, error) {
	secrets, err := secretManager.GetSecret()
	if err != nil {
		return nil, err
	}

	return &ConfigService{Secrets: secrets}, nil
}

func (c *ConfigService) Get(key string, defaultValue string) string {
	if value, exists := c.Secrets[key]; exists {
		return value
	}
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	if defaultValue == "" {
		log.Fatalf("Configuration key %s not found and no default value provided", key)
	}
	return defaultValue
}
