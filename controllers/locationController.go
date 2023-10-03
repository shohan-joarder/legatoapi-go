package controllers

import (
	"fmt"
	"net/http"

	"api.legatodesigns.com/database"
	"api.legatodesigns.com/models"
	"github.com/gin-gonic/gin"
)

type Country struct{
	ID int
	Name string
}

func GetCountry(c *gin.Context) {

	err := database.DBConnect()

	fmt.Println(err);

	list := models.CountryList();

	c.JSON(http.StatusOK, list)

	// return arr[0]
}



func GetCity(c *gin.Context) {

	// prams := c.Param("country_id")

	// println(prams)

	arr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}

	c.JSON(http.StatusOK, arr)

	// return arr[0]
}
func GetArea(c *gin.Context) {
	arr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}

	c.JSON(http.StatusOK, arr)

	// return arr[0]
}
