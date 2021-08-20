package service

import "github.com/ra192/card-api-go/model"

const CashAccountId = 1

func GetActiveAccountById(id uint) (model.Account, error) {
	Account := model.Account{}
	result := model.DB.First(&Account, model.Account{ID: id, Active: true})
	return Account, result.Error
}

func GetCashAccount() (model.Account, error) {
	return GetActiveAccountById(CashAccountId)
}
