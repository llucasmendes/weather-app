package weather

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetWeather(c *gin.Context) {
	zipcode := c.Param("zipcode")
	if len(zipcode) != 8 {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"message": "invalid zipcode"})
		return
	}

	city, err := getCityFromZipcode(zipcode)
	if err != nil {
		log.Printf("Error getting city from zipcode: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"message": "can not find zipcode"})
		return
	}

	tempC, err := getTemperature(city)
	if err != nil {
		log.Printf("Error getting temperature: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error getting temperature"})
		return
	}

	tempF := tempC*1.8 + 32
	tempK := tempC + 273

	c.JSON(http.StatusOK, gin.H{
		"temp_C": tempC,
		"temp_F": tempF,
		"temp_K": tempK,
	})
}
