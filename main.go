package main

import (
	"fmt"
	"net/http"
	"os"

	"api.legatodesigns.com/controllers"
	"api.legatodesigns.com/database"
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
			database.InitDB()
			query := "SELECT id, name,iso_code_2,iso_code_3,iso_numeric_code,address_format,postcode_required,phonecode,ordering,status FROM country"
			rows, err := database.DB.Query(query)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
				return
			}
			defer rows.Close()
			fmt.Println(rows)
			// Create a slice to store todos

			response := []Country{}
			// Iterate through the query results and populate the todos slice
			for rows.Next() {
				var country Country
				err := rows.Scan(&country.ID, &country.Name, &country.IsoCode2, &country.IsoCode3, &country.IsoNumberCode, &country.AddressFormat, &country.PostCodeRequired, &country.PhoneCode, &country.Ordering, &country.Status)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
					return
				}
				response = append(response, country)
			}

			// Return JSON response
			c.JSON(http.StatusOK, gin.H{"success": true, "message": "Country found", "data": response})

			// c.JSON(http.StatusOK, "Started api server")
		})
		api.GET("/country", controllers.GetCountry)
		api.GET("/city/:country_id", controllers.GetCity)
		api.GET("/area/:city_id", controllers.GetArea)
	}

	router.Run(":" + port)

}
