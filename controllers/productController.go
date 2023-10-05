package controllers

import (
	"net/http"
	"strconv"

	"api.legatodesigns.com/database"
	"api.legatodesigns.com/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	categoryId := c.Query("category_id")
	subCategoryId := c.Query("subcategory_id")

	pageStr := c.DefaultQuery("page", "1")
	limitStr := c.DefaultQuery("limit", "10")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	// limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Error getting query prams", "err": err.Error()})
	// 	return
	// }
	// limit := c.Query("limit")
	// perPage := c.Query("per_page")
	// currentPate := c.Query("current_pate")

	/* 	query := `
		SELECT p.id, p.name, p.slug, p.thumbnail_img, cat.name as cat_name, s_cat.name as sub_cat_name FROM products p
		LEFT JOIN categories c ON cat.id = p.category_id
		LEFT JOIN  sub_categories s_cat ON s_cat.id = p.subcategory_id
	` */

	query := `
		SELECT p.id, p.name, p.slug, p.thumbnail_img FROM products p WHERE published = 1
	`

	if categoryId != "" {
		query += `AND category_id =` + categoryId
	}

	if subCategoryId != "" {
		query += `AND subcategory_id =` + subCategoryId
	}

	offset := (page - 1) * limit

	strOffset := strconv.Itoa(offset)

	query += ` LIMIT ` + strOffset + "," + limitStr

	database.InitDB()

	rows, err := database.DB.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}

	defer rows.Close()

	getData := []models.GetProductWithOnlyThumbnail{}

	// var totalCount int

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying total records", "err": err.Error()})
		return
	}

	for rows.Next() {
		getProductDataObj := &models.GetProductWithOnlyThumbnail{}

		// totalCount = rows.Scan(&totalCount)
		err := rows.Scan(&getProductDataObj.ID, &getProductDataObj.Name, &getProductDataObj.Slug, &getProductDataObj.ThumbnailImage)

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning data", "err": err.Error()})
			return
		}

		getData = append(getData, *getProductDataObj)
	}

	rows.NextResultSet()

	// Read the total count
	var totalCount int
	if rows.Next() {
		err := rows.Scan(&totalCount)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error", "err": err.Error()})
			return
		}
	}

	totalPages := totalCount / limit
	if totalCount%limit != 0 {
		totalPages++
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Product found", "data": getData, "total_page": totalPages})
}
