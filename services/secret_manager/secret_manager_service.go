package secret_manager

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
)

// SecretManagerService es la estructura que administra los secretos de AWS
type SecretManagerService struct {
	SecretName string
	RegionName string
}

var instance *SecretManagerService
var once sync.Once

// NewSecretManagerService crea una instancia única de SecretManagerService (Singleton)
func NewSecretManagerService(secretName, regionName string) *SecretManagerService {
	once.Do(func() {
		instance = &SecretManagerService{
			SecretName: secretName,
			RegionName: regionName,
		}
	})
	return instance
}

// GetSecret obtiene y deserializa el secreto desde AWS Secrets Manager
func (s *SecretManagerService) GetSecret() (map[string]string, error) {
	// Crear una nueva sesión de AWS usando la región proporcionada
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(s.RegionName)},
	)
	if err != nil {
		return nil, fmt.Errorf("failed to create AWS session: %w", err)
	}

	// Crear un nuevo cliente de Secrets Manager
	svc := secretsmanager.New(sess)
	input := &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(s.SecretName),
	}

	// Obtener el valor del secreto desde AWS Secrets Manager
	result, err := svc.GetSecretValue(input)
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve secret: %w", err)
	}

	// Deserializar el secreto JSON
	var secretString string
	if result.SecretString != nil {
		secretString = *result.SecretString
	} else {
		return nil, fmt.Errorf("secret is not in string format")
	}

	// Convertir el secreto JSON en un mapa de cadenas
	var secrets map[string]string
	err = json.Unmarshal([]byte(secretString), &secrets)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal secret: %w", err)
	}

	return secrets, nil
}
