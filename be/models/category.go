package models

type Category struct {
	ID    uint   `json:"id" gorm:"primary+key"`
	Title string `json:"title"`
	Count int32  `json:"count"`
}
