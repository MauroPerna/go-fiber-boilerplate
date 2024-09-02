package user

import (
	"context"
	"log"

	"github.com/MauroPerna/go-fiber-backend-test/services/database"
	entities "github.com/MauroPerna/go-fiber-backend-test/services/database/entities/domain"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
)

type UserService struct {
	db *database.DatabaseModule
}

func NewUserService(db *database.DatabaseModule) *UserService {
	return &UserService{db: db}
}

func (s *UserService) GetAllUsers(c *fiber.Ctx) error {
	collection := s.db.DatabaseService.GetCollection("APP", "owners")

	var owners []entities.Owner

	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		log.Println("Error fetching owners:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not fetch owners",
			"error":   err.Error(),
		})
	}
	defer cursor.Close(context.Background())

	if err := cursor.All(context.Background(), &owners); err != nil {
		log.Println("Error decoding owners:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"status":  "error",
			"message": "Could not decode owners",
			"error":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Fetched all owners",
		"data":    owners,
	})
}

func (s *UserService) GetUser(c *fiber.Ctx) error {
	id := c.Params("id")
	return c.JSON(fiber.Map{
		"status":  "success",
		"message": "Get user by ID",
		"id":      id,
	})
}
