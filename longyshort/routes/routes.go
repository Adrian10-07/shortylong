package routes

import (
	"shortYlong/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.POST("/personas", controllers.CreatePersona)
	r.GET("/personas", controllers.GetAllPersonas)
	r.GET("/personas/recent", controllers.GetRecentlyAddedPersonas) 
	r.GET("/personas/gender_count", controllers.GetGenderCount)     
}
