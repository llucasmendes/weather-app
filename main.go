package main

import (
	"weather-app/weather"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/weather/:zipcode", weather.GetWeather)

	port := "8080"
	r.Run(":" + port)
}
