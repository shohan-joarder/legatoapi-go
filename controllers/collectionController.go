package controllers

import (
	"database/sql"
	"net/http"

	"api.legatodesigns.com/database"
	"api.legatodesigns.com/helper"
	"api.legatodesigns.com/models"
	"github.com/gin-gonic/gin"
)

func GetCollection(c *gin.Context) {
	query :="SELECT title,slug, banner from collections"
	database.InitDB()
	rows, err := database.DB.Query(query)

	if err !=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}

	defer rows.Close()
	response :=[] models.Collection{}
	for rows.Next(){
		var collection models.Collection
		err :=rows.Scan(&collection.Title,&collection.Slug,&collection.Banner)
		if err !=nil{
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
		}
		if collection.Banner != ""{
			collection.Banner = helper.GetFilePath( collection.Banner)
		}
		response = append(response,collection)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Collection found", "data": response})
}


func getNextCollectionId(currentID int)int{
	var nextID int
	database.InitDB()
	err := database.DB.QueryRow("SELECT id FROM collections WHERE id > ? ORDER BY id ASC LIMIT 1", currentID).Scan(&nextID)

	if err != nil {
		if err == sql.ErrNoRows {

			var firstRowID int
			err = database.DB.QueryRow("SELECT id FROM collections ORDER BY id ASC LIMIT 1").Scan(&firstRowID)
			if err != nil {
				return 0
			}
			return firstRowID
		}
		return 0
	}
	return nextID
}

func GetCollectionDetails(c *gin.Context)  {
	pram :=c.Param("slug")
	database.InitDB()

	response  := models.CollectionDetails{}

	err := database.DB.QueryRow("SELECT id,title,slug,subtitle,about_title,about_description,	banner,	og_image,	meta_title,	meta_keywords,	meta_description,	og_title,	og_description from collections where slug = ?", pram).Scan(&response.ID,&response.Title,&response.Slug,&response.SubTitle,&response.AboutTitle,&response.AboutDescription,&response.Banner,&response.OgImage,&response.MetaTitle,&response.MetaKeywords,&response.MetaDescription,&response.OgTitle,&response.OgDescription)

	if response.Banner != ""{
		response.Banner = helper.GetFilePath( response.Banner)
	}

	nextId := getNextCollectionId(response.ID)

	if nextId !=0{
		var collection models.Collection
		err := database.DB.QueryRow("SELECT title,slug, banner from collections where id = ?", nextId).Scan(&collection.Title,&collection.Slug,&collection.Banner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows next collection", "err": err.Error()})
				return
		}

		if collection.Banner != ""{
			collection.Banner = helper.GetFilePath( collection.Banner)
		}
	
		response.NextCollection = collection
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Collection found", "data": response})
}