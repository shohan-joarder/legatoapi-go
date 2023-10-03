package models

type Collection struct {
	Title  string `json:"title"`
	Slug   string `json:"slug"`
	Banner string `json:"banner"`
}

type CollectionDetails struct {
	ID               int        `json:"id"`
	Title            string     `json:"title"`
	Slug             string     `json:"slug"`
	SubTitle         string     `json:"subtitle"`
	AboutTitle       *string    `json:"about_title"`
	AboutDescription *string    `json:"about_description"`
	Banner           string     `json:"banner"`
	OgImage          *string    `json:"og_image"`
	MetaTitle        *string    `json:"meta_title"`
	MetaKeywords     *string    `json:"meta_keywords"`
	MetaDescription  *string    `json:"meta_description"`
	OgTitle          *string    `json:"og_title"`
	OgDescription    *string    `json:"og_description"`
	NextCollection   Collection `json:"next_collection"`
}