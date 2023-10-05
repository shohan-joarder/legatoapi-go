package models

import "database/sql"

type Collection struct {
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Banner string `json:"banner"`
}

type CollectionCategory struct {
	ID int `json:"id"`
	Name string `json:"cat_name"`
	Slug string `json:"cat_slug"`
	Products [] ResProductWithOnlyThumbnail `json:"products"`
}

type GetCollectionDetails struct{
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
	NextCollection   Collection     `json:"next_collection"`
	CollectionCategories CollectionCategory `json:"collection_categories"`
}

type ResponseCollectionDetails struct {
	ID               int            `json:"id"`
	Title            string         `json:"title"`
	Slug             string         `json:"slug"`
	SubTitle         string         `json:"subtitle"`
	AboutTitle       *string        `json:"about_title"`
	AboutDescription *string        `json:"about_description"`
	Banner           string         `json:"banner"`
	OgImage          string 		`json:"og_image,omitempty"`
	MetaTitle        *string        `json:"meta_title"`
	MetaKeywords     *string        `json:"meta_keywords"`
	MetaDescription  *string        `json:"meta_description"`
	OgTitle          *string        `json:"og_title"`
	OgDescription    *string        `json:"og_description"`
	NextCollection   Collection     `json:"next_collection"`
	CollectionCategories CollectionCategory `json:"collection_categories"`
}

type GetCollection struct {
	ID      int            `json:"id"`
	Title   string         `json:"title"`
	Banner  sql.NullString `json:"banner,omitempty"`
	OgImage sql.NullString `json:"og_image,omitempty"`
}

type ResponseCollection struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Banner  string `json:"banner,omitempty"`
	OgImage string `json:"og_image,omitempty"`
}
