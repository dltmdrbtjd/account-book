package models

import "time"

type Account struct {
	ID       uint   `json:"id" gorm:"primary+key"`
	Title    string `json:"title"`
	Category string `json:"category"`
	Money    int32  `json:"money"`
	CreateAt time.Time
	UpdateAt time.Time
}
