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
	ID        int   
	Name     string 
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
		api.GET("/",func (c *gin.Context)  {
			database.InitDB()
			rows, err := database.DB.Query("SELECT * FROM country WHERE status=true")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database"})
				return
			}
			defer rows.Close()
			fmt.Println(rows)
		// Create a slice to store todos
		country := Country{}
		response :=[] Country{}
		// Iterate through the query results and populate the todos slice
		for rows.Next() {
			var id int
			var name string
			err := rows.Scan(&country.ID, &country.Name)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows"})
				return
			}
			country.ID = id
			country.Name = name
			response = append(response, country)
		}

		// Return JSON response
		c.JSON(http.StatusOK, response)

			// c.JSON(http.StatusOK, "Started api server")
		})
		api.GET("/country", controllers.GetCountry)
		api.GET("/city/:country_id", controllers.GetCity)
		api.GET("/area/:city_id", controllers.GetArea)
	}

	router.Run(":" + port)

}
