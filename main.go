package main

import (
	"fmt"
	"net/http"
	"os"

	"api.legatodesigns.com/controllers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type Country struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	IsoCode2         string `json:"iso_code_2"`
	IsoCode3         string `json:"iso_code_3"`
	IsoNumberCode    int    `json:"iso_numeric_code"`
	AddressFormat    string `json:"address_format"`
	PostCodeRequired int8   `json:"postcode_required"`
	PhoneCode        int    `json:"phonecode"`
	Ordering         int    `json:"ordering"`
	Status           int8   `json:"status"`
}

func main() {
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
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Welcome to api routing"})
		})
		api.GET("/get-country", controllers.GetCountry)
		api.GET("/get-states/:country_id", controllers.GetState)
		api.GET("/get-city/:state_id", controllers.GetCity)
		api.GET("/get-area/:city_id", controllers.GetArea)

		// collection routes
		api.GET("/collections",controllers.GetCollection)
		api.GET("/collections/:slug",controllers.GetCollectionDetails)
	}

	router.Run(":" + port)

}
