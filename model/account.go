package model

type Account struct {
	ID uint `gorm:"primary_key"`
	Name string
	Active bool
	Currency string
	MerchantID uint
	Merchant Merchant
}
