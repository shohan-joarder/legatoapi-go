package controllers

import (
	"net/http"

	"api.legatodesigns.com/database"
	"api.legatodesigns.com/models"
	"github.com/gin-gonic/gin"
)

func GetCountry(c *gin.Context) {

	query := "SELECT id, name,iso_code_2,iso_code_3,iso_numeric_code,address_format,postcode_required,phonecode,ordering,status FROM country where status=1"
	database.InitDB()
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}
	defer rows.Close()

	response := []models.Country{}
	for rows.Next() {
		var country models.Country
		err := rows.Scan(&country.ID, &country.Name, &country.IsoCode2, &country.IsoCode3, &country.IsoNumberCode, &country.AddressFormat, &country.PostCodeRequired, &country.PhoneCode, &country.Ordering, &country.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
		}
		response = append(response, country)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Country found", "data": response})
}

func GetState(c *gin.Context) {
	pram := c.Param("country_id")
	query := "SELECT id,name FROM state where country_id=" + pram
	database.InitDB()
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}
	defer rows.Close()

	response := []models.State{}
	for rows.Next() {
		var state models.State
		err := rows.Scan(&state.Id, &state.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
		}
		response = append(response, state)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "State found", "data": response})
}

func GetCity(c *gin.Context) {

	pram := c.Param("state_id")
	query := "SELECT id,title FROM locations where state_id=" + pram + " AND  parent_id IS NULL"
	database.InitDB()
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}
	defer rows.Close()

	response := []models.CityAndLocation{}
	for rows.Next() {
		var city models.CityAndLocation
		err := rows.Scan(&city.Id, &city.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
		}
		response = append(response, city)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "City found", "data": response})
}
func GetArea(c *gin.Context) {
	pram := c.Param("city_id")
	query := "SELECT id,title FROM locations where parent_id=" + pram
	database.InitDB()
	rows, err := database.DB.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}
	defer rows.Close()

	response := []models.CityAndLocation{}
	for rows.Next() {
		var city models.CityAndLocation
		err := rows.Scan(&city.Id, &city.Title)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
		}
		response = append(response, city)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Area found", "data": response})
}
