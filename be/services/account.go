package services

import (
	"api/models"
	"time"
)

type CreateAccountInput struct {
	Title    string    `json:"title" binding:"required"`
	Category string    `json:"category" binding:"required"`
	Money    int32     `json:"money" binding:"required"`
	CreateAt time.Time `json:"create_at" binding:"required"`
}

type UpdateAccountInput struct {
	Title    string    `json:"title"`
	Category string    `json:"category"`
	Money    int32     `json:"money"`
	UpdateAt time.Time `json:"create_at" binding:"required"`
}

func FindAccounts() []models.Account {
	var accounts []models.Account
	models.DB.Find(&accounts)

	return accounts
}

func CreateAccount() CreateAccountInput {
	var input CreateAccountInput
	input.CreateAt = time.Now()

	return input
}

func UpdateAccount() UpdateAccountInput {
	var input UpdateAccountInput
	input.UpdateAt = time.Now()

	return input
}
