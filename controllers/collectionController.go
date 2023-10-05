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
	query := "SELECT title,slug, banner from collections"
	database.InitDB()
	rows, err := database.DB.Query(query)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error querying database", "err": err.Error()})
		return
	}

	defer rows.Close()
	response := []models.Collection{}
	for rows.Next() {
		var collection models.Collection
		err := rows.Scan(&collection.Title, &collection.Slug, &collection.Banner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
			return
		}
		if collection.Banner != "" {
			collection.Banner = helper.GetFilePath(collection.Banner)
		}
		response = append(response, collection)
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Collection found", "data": response})
}

func getNextCollectionId(currentID int) int {
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

type GetCollectionDetailsST struct{
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	Slug             string         `json:"slug"`
	SubTitle         string         `json:"subtitle"`
	AboutTitle       *string        `json:"about_title"`
	AboutDescription *string        `json:"about_description"`
	Banner           sql.NullString `json:"banner"`
	OgImage          sql.NullString `json:"og_image,omitempty"`
	MetaTitle        *string        `json:"meta_title"`
	MetaKeywords     *string        `json:"meta_keywords"`
	MetaDescription  *string        `json:"meta_description"`
	OgTitle          *string        `json:"og_title"`
	OgDescription    *string        `json:"og_description"`
	NextCollection   models.Collection     `json:"next_collection"`
	// CollectionCategories []models.CollectionCategory `json:"collection_categories"`
}
type ResponseCollectionDetailsST struct{
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	Slug             string         `json:"slug"`
	SubTitle         string         `json:"subtitle"`
	AboutTitle       *string        `json:"about_title"`
	AboutDescription *string        `json:"about_description"`
	Banner           string 		`json:"banner"`
	OgImage          string 		`json:"og_image,omitempty"`
	MetaTitle        *string        `json:"meta_title"`
	MetaKeywords     *string        `json:"meta_keywords"`
	MetaDescription  *string        `json:"meta_description"`
	OgTitle          *string        `json:"og_title"`
	OgDescription    *string        `json:"og_description"`
	NextCollection   models.Collection     `json:"next_collection"`
	CollectionCategories []models.CollectionCategory `json:"collection_categories"`
}


func getCollectionBySlug(slug string)(GetCollectionDetailsST,error){
	database.InitDB()
	getCollNullSql :=GetCollectionDetailsST{}
	err := database.DB.QueryRow("SELECT id,title,slug,subtitle,about_title,about_description,	banner,	og_image,	meta_title,	meta_keywords,	meta_description,	og_title,	og_description from collections where slug = ?", slug).Scan(&getCollNullSql.ID, &getCollNullSql.Title, &getCollNullSql.Slug, &getCollNullSql.SubTitle, &getCollNullSql.AboutTitle, &getCollNullSql.AboutDescription, &getCollNullSql.Banner, &getCollNullSql.OgImage, &getCollNullSql.MetaTitle, &getCollNullSql.MetaKeywords, &getCollNullSql.MetaDescription, &getCollNullSql.OgTitle, &getCollNullSql.OgDescription)

	if err != nil {
		return getCollNullSql, err
	}
	return getCollNullSql, nil

}


func getCollectionCategoryProducts(id int) ([]models.ResProductWithOnlyThumbnail,error){
	database.InitDB()
	

	query :=`
		SELECT p.id, p.name,p.slug,p.thumbnail_img FROM products p 
		LEFT JOIN collection_category_product ccp ON ccp.product_id = p.id 
		WHERE ccp.collection_category_id=?
	` 

	rows, err := database.DB.Query(query,id)

	if err != nil {
		return []models.ResProductWithOnlyThumbnail{}, err
	}
	defer rows.Close()

	products :=[]models.ResProductWithOnlyThumbnail{}

	for rows.Next() {
		getCollCategoryProduct :=models.GetProductWithOnlyThumbnail{}
		resCollCategoryProduct :=models.ResProductWithOnlyThumbnail{}
		err := rows.Scan(&getCollCategoryProduct.ID,&getCollCategoryProduct.Name, &getCollCategoryProduct.Slug,&getCollCategoryProduct.ThumbnailImage)

		if err != nil {
			return []models.ResProductWithOnlyThumbnail{}, err
		}

		resCollCategoryProduct.ID = getCollCategoryProduct.ID
		resCollCategoryProduct.Name = getCollCategoryProduct.Name
		resCollCategoryProduct.Slug = getCollCategoryProduct.Slug

		if getCollCategoryProduct.ThumbnailImage.Valid{
			resCollCategoryProduct.ThumbnailImage = helper.GetFilePath(getCollCategoryProduct.ThumbnailImage.String)
		}

		products = append(products, resCollCategoryProduct)
	}
	return products, nil
}

func getCollectionCategoryById(id int)([]models.CollectionCategory,error){
	database.InitDB()
	
	returnData :=[]models.CollectionCategory{}

	query :=`
		SELECT cc.id, cc.name,cc.slug FROM collection_categories cc
		LEFT JOIN collection_collection_category  ccc ON ccc.collection_id = cc.id
		Where ccc.collection_id =?
	` 

	rows, err := database.DB.Query(query,id)

	if err != nil {
		return []models.CollectionCategory{}, err
	}

	defer rows.Close()

	for rows.Next() {
		getCollCategory :=models.CollectionCategory{}
		err := rows.Scan(&getCollCategory.ID,&getCollCategory.Name, &getCollCategory.Slug)

		if err != nil {
			return []models.CollectionCategory{}, err
		}

		categoryProduct,err := getCollectionCategoryProducts(getCollCategory.ID)

		if err != nil {
			return []models.CollectionCategory{}, err
		}

		getCollCategory.Products= categoryProduct

		returnData = append(returnData, getCollCategory)
	}

	return returnData , nil
}

func GetCollectionDetails(c *gin.Context) {
	pram := c.Param("slug")
	
	collection ,err :=getCollectionBySlug(pram)

	if err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error get collection by slug", "err": err.Error()})
			return
	}

	response := ResponseCollectionDetailsST{}

	response.ID = collection.ID
	response.Title = collection.Title
	response.Slug = collection.Slug
	response.SubTitle = collection.SubTitle
	response.AboutTitle = collection.AboutTitle
	response.AboutDescription = collection.AboutTitle
	response.MetaTitle = collection.MetaTitle
	response.MetaKeywords = collection.MetaKeywords
	response.MetaDescription = collection.MetaDescription
	response.OgTitle = collection.OgTitle
	response.OgDescription = collection.OgDescription

	if collection.Banner.Valid {
		response.Banner = helper.GetFilePath(collection.Banner.String)
	}else{
		response.Banner = ""
	}
	if collection.OgImage.Valid {
		response.OgImage = helper.GetFilePath(collection.OgImage.String)
	}else{
		response.OgImage = ""
	}

	response.CollectionCategories,err= getCollectionCategoryById(response.ID)

	nextId := getNextCollectionId(response.ID)

	if nextId != 0 {
		var collection models.Collection
		err := database.DB.QueryRow("SELECT title,slug, banner from collections where id = ?", nextId).Scan(&collection.Title, &collection.Slug, &collection.Banner)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows next collection", "err": err.Error()})
			return
		}

		if collection.Banner != "" {
			collection.Banner = helper.GetFilePath(collection.Banner)
		}

		response.NextCollection = collection
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Collection found", "data": response})
}

type GetCollectionST struct {
	ID      int            `json:"id"`
	Title   string         `json:"title"`
	Banner  sql.NullString `json:"banner,omitempty"`
	OgImage sql.NullString `json:"og_image,omitempty"`
}

type ResponseCollectionST struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Banner  string `json:"banner,omitempty"`
	OgImage string `json:"og_image,omitempty"`
}

func getCollectionFromDB(collectionId int) (GetCollectionST, error) {
	var collection GetCollectionST
	database.InitDB()

	err := database.DB.QueryRow("SELECT id, title,banner,og_image FROM collections WHERE id=?", collectionId).Scan(&collection.ID, &collection.Title, &collection.Banner, &collection.OgImage)

	if err != nil {
		return collection, err
	}

	return collection, nil
}

func HandleGetCollectionAPIRequest(c *gin.Context) {
	collection, err := getCollectionFromDB(1) // Example user ID

	if err != nil {
		// Handle database error (e.g., log the error, return an error response)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows next collection", "err": err.Error()})
		return
	}

	apiResponse := ResponseCollectionST{
		ID:     collection.ID,
		Title:  collection.Title,
		Banner: collection.Banner.String,
	}

	// Check if the Name field is valid (not null)
	if collection.OgImage.Valid {
		// Concatenate a string with the Name value
		apiResponse.OgImage = collection.OgImage.String
		apiResponse.OgImage = helper.GetFilePath(apiResponse.OgImage)
	} else {
		// Handle the case when the Name field is null
		apiResponse.OgImage = "Hello, Guest!"
	}

	// Convert API response to JSON and send it
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error scanning rows", "err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "message": "Collection found", "data": apiResponse})
}


/* func getCollectionWithCategory(catSlug string)([]GetCollectionDetails,error){
	database.InitDB()
		rows, err := database.DB.Query(`
		SELECT col.id, col.title, col.slug,col.subtitle,col.about_title,col.about_description,col.banner,col.og_image,
		col.meta_title,col.meta_keywords,col.meta_description,col.og_title,col.og_description,cat.slug,cat.name
		FROM collections col
		LEFT JOIN collection_collection_category pccc ON col.id = pccc.collection_id
		LEFT JOIN collection_categories cat ON cat.id = pccc.collection_category_id
		WHERE col.slug=?`,catSlug)

		if err != nil {
			return nil, err
		}
		defer rows.Close()

		collections := make(map[int]GetCollectionDetails)

		resCollection := GetCollectionDetails{}
		resCat :=models.CollectionCategory{}

		for rows.Next() {
			err := rows.Scan(&resCollection.ID, &resCollection.Title, &resCollection.Slug,&resCollection.Slug, &resCollection.SubTitle, &resCollection.AboutTitle, &resCollection.AboutDescription, &resCollection.Banner, &resCollection.OgImage, &resCollection.MetaTitle, &resCollection.MetaKeywords, &resCollection.MetaDescription, &resCollection.OgTitle, &resCollection.OgDescription, &resCollection,resCat.Slug,resCat.Name)
			if err != nil {
				return nil, err
			}

		}

} */