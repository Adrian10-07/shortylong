package routes

import (
	"shortYlong/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
)

func SetupRoutes(r *gin.Engine) {
	// Configuración básica de CORS
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // O especifica tus orígenes: {"http://127.0.0.1:5501"}
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	r.POST("/personas", controllers.CreatePersona)
	r.GET("/personas", controllers.GetAllPersonas)
	r.GET("/personas/recent", controllers.GetRecentlyAddedPersonas) 
	r.GET("/personas/gender_count", controllers.GetGenderCount)
}