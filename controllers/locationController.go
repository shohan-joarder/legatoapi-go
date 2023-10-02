package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCountry(c *gin.Context) {
	arr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}

	c.JSON(http.StatusOK, arr)

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
