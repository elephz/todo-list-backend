package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	BirthDate time.Time      `json:"birth_date"`
	CreatedAt time.Time      `json:"created_at"`
	Address   string         `json:"address"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}

var db *gorm.DB

func initDatabase() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	user := os.Getenv("DB_USER")
	pass := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user, pass, host, port, name)

	database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&User{})
	db = database
	fmt.Println("Database connected")
}

func main() {
	app := fiber.New()

	// Enable CORS globally
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Allow all origins (you can specify more specific origins if needed)
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	initDatabase()

	api := app.Group("/api")
	api.Get("/users", getUsers)
	api.Get("/users/:id", getUser)
	api.Post("/users", createUser)

	log.Fatal(app.Listen(":3000"))
}

func getUsers(c *fiber.Ctx) error {
	var users []User
	db.Find(&users)
	return c.JSON(users)
}

func getUser(c *fiber.Ctx) error {
	id := c.Params("id")
	var user User
	if err := db.First(&user, id).Error; err != nil {
		return c.Status(404).SendString("User not found")
	}
	return c.JSON(user)
}

func createUser(c *fiber.Ctx) error {
	user := new(User)
	if err := c.BodyParser(user); err != nil {
		return c.Status(400).SendString(err.Error())
	}
	db.Create(&user)
	return c.Status(201).JSON(user)
}
