package database

import (
	"context"
	"log"
	"time"

	"github.com/MauroPerna/go-fiber-backend-test/services/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/fx"
)

type DatabaseService struct {
	AppClient    *mongo.Client
	SystemClient *mongo.Client
}

// NewDatabaseService se encarga de crear las conexiones necesarias a las bases de datos.
func NewDatabaseService(lc fx.Lifecycle, c *config.ConfigService) *DatabaseService {
	service := &DatabaseService{}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Conexión a la base de datos principal (APP)
			appClient, err := connectToDatabase(ctx, c.Get("MONGODB_READ_WRITE_URI", ""))
			if err != nil {
				return err
			}
			service.AppClient = appClient

			// Conexión a la base de datos del sistema (SYSTEM)
			systemClient, err := connectToDatabase(ctx, c.Get("MONGODB_READ_ONLY_URI", ""))
			if err != nil {
				return err
			}
			service.SystemClient = systemClient

			log.Println("Connected to MongoDB databases!")
			return nil
		},
		OnStop: func(ctx context.Context) error {
			disconnectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if service.AppClient != nil {
				if err := service.AppClient.Disconnect(disconnectCtx); err != nil {
					return err
				}
			}
			if service.SystemClient != nil {
				if err := service.SystemClient.Disconnect(disconnectCtx); err != nil {
					return err
				}
			}

			log.Println("Disconnected from MongoDB databases.")
			return nil
		},
	})

	return service
}

// connectToDatabase es una función auxiliar para conectar a una base de datos MongoDB utilizando la URI proporcionada.
func connectToDatabase(ctx context.Context, uri string) (*mongo.Client, error) {
	connectCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(connectCtx, clientOptions)
	if err != nil {
		return nil, err
	}

	if err = client.Ping(connectCtx, nil); err != nil {
		return nil, err
	}

	return client, nil
}

// GetCollection es una función para obtener una colección específica de una base de datos dada.
func (d *DatabaseService) GetCollection(database, collectionName string) *mongo.Collection {
	switch database {
	case "APP":
		return d.AppClient.Database("app").Collection(collectionName)
	case "SYSTEM":
		return d.SystemClient.Database("system").Collection(collectionName)
	default:
		log.Fatalf("Invalid database name provided: %s", database)
		return nil
	}
}
