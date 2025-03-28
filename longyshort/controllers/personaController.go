package controllers

import (
	"shortYlong/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func CreatePersona(c *gin.Context) {
	var persona models.Persona

	if err := c.ShouldBindJSON(&persona); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.CreatePersona(&persona); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}


	c.JSON(http.StatusCreated, gin.H{"message": "Persona creada exitosamente", "persona": persona})
}

func GetAllPersonas(c *gin.Context) {
	personas, err := models.GetAllPersonas()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personas)
}

// Short Polling: Obtener las personas recientemente añadidas
func GetRecentlyAddedPersonas(c *gin.Context) {
	personas, err := models.GetRecentPersonas(10) 
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, personas)
}

// Long Polling: Contador de géneros
func GetGenderCount(c *gin.Context) {

	for {
		count, err := models.GetGenderCount() 
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if count != nil {
			c.JSON(http.StatusOK, count)
			return
		}

		time.Sleep(5 * time.Second) // Esperar 5 segundos antes de verificar nuevamente.
	}
}
