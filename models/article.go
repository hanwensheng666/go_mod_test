package models

type Article struct {
	Model
	TagID int    `json:""tag_id gorm:"index"`
	Tag   Tag    `json:"tag"`
	Title string `json:"title"`
}
