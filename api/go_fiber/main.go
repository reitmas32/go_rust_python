package main

import (
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID   int    `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Abrir la conexión con la base de datos SQLite
	db, err := gorm.Open(sqlite.Open("users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Error al abrir la base de datos:", err)
	}
	// Migrar el esquema de la base de datos
	db.AutoMigrate(&User{})

	// Crear una instancia de Fiber
	app := fiber.New()

	// Definir el endpoint GET /users/:id
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		// Obtener el ID del parámetro de la ruta
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid ID"})
		}
		// Consultar el usuario por su ID
		var user User
		result := db.First(&user, id)
		if result.Error != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "User not found"})
		}
		// Devolver el usuario como JSON
		return c.JSON(user)
	})

	// Iniciar el servidor Fiber en el puerto 8000
	log.Fatal(app.Listen(":8000"))
}
