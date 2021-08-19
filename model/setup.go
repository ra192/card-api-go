package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dsn = "host=localhost user=card-api password=card-api dbname=card-api-go"
var DB *gorm.DB

func ConnectDataBase() {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err1 := db.AutoMigrate(&Merchant{}, &Account{}, &Customer{}, &Card{}, &Transaction{}, &TransactionItem{},
		&TransactionFee{})
	if err1 != nil {
		panic("failed to migrate database")
	}

	DB = db
}
