package auth

import (
	"context"
	"fmt"
	"log"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/MauroPerna/go-fiber-backend-test/services/config"
	"google.golang.org/api/option"
)

type AuthService struct {
	AppName     string
	FirebaseApp *firebase.App
}

func NewAuthService(configService *config.ConfigService) *AuthService {
	currentDirectory, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to get current directory: %v", err)
	}

	credPath := fmt.Sprintf("%s/firebaseSDKconfig.json", currentDirectory)
	cred := option.WithCredentialsFile(credPath)
	appName := configService.Get("FIREBASE_APP_NAME", "default")

	// Initialize Firebase App
	firebaseApp, err := firebase.NewApp(context.Background(), nil, cred)
	if err != nil {
		log.Fatalf("Failed to initialize Firebase app: %v", err)
	}

	return &AuthService{
		AppName:     appName,
		FirebaseApp: firebaseApp,
	}
}

func (a *AuthService) VerifyAccessToken(token string) (*auth.Token, error) {
	client, err := a.FirebaseApp.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("failed to get Auth client: %v", err)
	}

	decodedToken, err := client.VerifyIDToken(context.Background(), token)
	if err != nil {
		return nil, fmt.Errorf("failed to verify ID token: %v", err)
	}

	return decodedToken, nil
}
