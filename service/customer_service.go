package service

import (
	"github.com/ra192/card-api-go/model"
	"log"
)

func CreateCustomer(customer *model.Customer)  {
	model.DB.Create(customer)
	log.Printf("Customer was created with id: %d",customer.ID)
}