package service

import "github.com/ra192/card-api-go/model"

func GetMerchantById(ID uint) (model.Merchant, error) {
	Merchant := model.Merchant{}
	result := model.DB.First(&Merchant, ID)

	return Merchant, result.Error
}
