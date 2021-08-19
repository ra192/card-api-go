package model

import (
	"gorm.io/datatypes"
)

type Customer struct {
	ID uint `gorm:"primary_key"`
	Phone	string
	Email string
	Active bool
	FirstName string
	LastName string
	BirthDate datatypes.Date
	Address string
	Address2 string
	City string
	StateRegion string
	Country string
	PostalCode string
	MerchantID uint
	Merchant Merchant
}
