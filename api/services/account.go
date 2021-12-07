package services

import (
	"api/models"
)

func FindAccounts() []models.Account {
	var accounts []models.Account
	models.DB.Find(&accounts)

	return accounts
}
