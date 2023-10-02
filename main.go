package main

import (
	"fmt"
	"os"

	"api.legatodesigns.com/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	connect := Dbconnect()

	if connect != nil {
		fmt.Println(connect)
	}

	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
	// load env here
	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "8080"
	}

	router := gin.Default()

	api := router.Group("/api")
	{
		api.GET("/country", controllers.GetCountry)
		api.GET("/city/:country_id", controllers.GetCity)
		api.GET("/area/:city_id", controllers.GetArea)
	}

	router.Run(":" + port)

}
