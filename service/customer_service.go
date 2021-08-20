package service

import (
	"github.com/ra192/card-api-go/model"
	"log"
)

func CreateCustomer(customer *model.Customer) error {
	result := model.DB.Create(customer)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Customer was created with id: %d", customer.ID)
	return nil
}
