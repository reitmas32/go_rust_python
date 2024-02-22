package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID      int
	Name    string
	Age     int
	Email   string
	City    string
	Country string
}

func main() {
	// Abrir la conexión con la base de datos SQLite
	db, err := sql.Open("sqlite3", "users.db")
	if err != nil {
		fmt.Println("Error al abrir la base de datos:", err)
		return
	}
	defer db.Close()

	// Realizar algunas consultas
	fmt.Println("Consultando algunos datos de usuarios:")

	// Consulta 1: Obtener todos los usuarios
	fmt.Println("\nTodos los usuarios:")
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.City, &user.Country)
		if err != nil {
			fmt.Println("Error al escanear fila:", err)
			return
		}
		//fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s, City: %s, Country: %s\n", user.ID, user.Name, user.Age, user.Email, user.City, user.Country)
	}

	// Consulta 2: Obtener usuarios mayores de 30 años
	fmt.Println("\nUsuarios mayores de 30 años:")
	rows, err = db.Query("SELECT * FROM users WHERE Age > ?", 30)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Email, &user.City, &user.Country)
		if err != nil {
			fmt.Println("Error al escanear fila:", err)
			return
		}
		//fmt.Printf("ID: %d, Name: %s, Age: %d, Email: %s, City: %s, Country: %s\n", user.ID, user.Name, user.Age, user.Email, user.City, user.Country)
	}
}
