package model

type Merchant struct {
	ID uint `gorm:"primary_key"`
	Name string
	Secret string
}