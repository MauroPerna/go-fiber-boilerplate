package secret_manager

import "go.uber.org/fx"

type SecretManagerModule struct {
	SecretManagerService *SecretManagerService
}

func NewSecretManagerModule() *SecretManagerModule {
	return &SecretManagerModule{
		SecretManagerService: NewSecretManagerService("self-hosted-demo-dev-envs", "us-east-1"),
	}
}

var Module = fx.Options(
	fx.Provide(
		NewSecretManagerModule,
		func(m *SecretManagerModule) *SecretManagerService { return m.SecretManagerService },
	),
)
