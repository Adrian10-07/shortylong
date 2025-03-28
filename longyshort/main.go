package main

import (
	"fmt"
	"shortYlong/routes"
	"shortYlong/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la conexión a la base de datos
    models.InitDB()


	// Crear el router de Gin
	r := gin.Default()

	// Configurar las rutas
	routes.SetupRoutes(r)

	// Iniciar el servidor
	if err := r.Run(":8080"); err != nil {
		fmt.Println("Error al iniciar el servidor: ", err)
	}
}
