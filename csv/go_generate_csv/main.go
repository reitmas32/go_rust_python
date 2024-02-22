package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// User representa la estructura de un usuario
type User struct {
	ID     int
	Nombre string
	Edad   int
	Email  string
	Ciudad string
	Pais   string
}

func main() {
	// Generar 1 millón de registros de usuarios
	cantidadRegistros := 10000
	usuarios := make([]User, cantidadRegistros)
	for i := 0; i < cantidadRegistros; i++ {
		usuarios[i] = User{
			ID:     i + 1,
			Nombre: "Usuario" + strconv.Itoa(i+1),
			Edad:   18 + (i % 50), // Edad entre 18 y 67 años
			Email:  "usuario" + strconv.Itoa(i+1) + "@example.com",
			Ciudad: "Ciudad" + strconv.Itoa(i%100), // 100 ciudades
			Pais:   "Pais" + strconv.Itoa(i%10),    // 10 países
		}
	}

	// Escribir los usuarios en un archivo CSV
	archivoCSV, err := os.Create("usuarios.csv")
	if err != nil {
		fmt.Println("Error al crear el archivo CSV:", err)
		return
	}
	defer archivoCSV.Close()

	escritorCSV := csv.NewWriter(archivoCSV)
	defer escritorCSV.Flush()

	// Escribir el encabezado
	encabezado := []string{"ID", "Nombre", "Edad", "Email", "Ciudad", "Pais"}
	escritorCSV.Write(encabezado)

	// Escribir los registros de usuarios
	for _, usuario := range usuarios {
		fila := []string{
			strconv.Itoa(usuario.ID),
			usuario.Nombre,
			strconv.Itoa(usuario.Edad),
			usuario.Email,
			usuario.Ciudad,
			usuario.Pais,
		}
		escritorCSV.Write(fila)
	}

	fmt.Println("Archivo CSV generado exitosamente.")
}
