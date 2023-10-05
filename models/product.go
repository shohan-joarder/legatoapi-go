package models

import "database/sql"

type GetProductWithOnlyThumbnail struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	ThumbnailImage sql.NullString `json:"thumbnail_img"`
}
type ResProductWithOnlyThumbnail struct {
	ID             int            `json:"id"`
	Name           string         `json:"name"`
	Slug           string         `json:"slug"`
	ThumbnailImage string `json:"thumbnail_img"`
}