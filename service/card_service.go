package service

import (
	"github.com/ra192/card-api-go/model"
	"log"
)

const (
	CardAccountId = 2
	FeeAccountId  = 3
)

func СreateCard(card *model.Card) error {
	result := model.DB.Create(card)
	if result.Error != nil {
		return result.Error
	}
	log.Printf("Customer was created with id: %d", card.ID)
	return nil
}

func DepositCard(cardId uint, amount uint, orderId string) (model.Transaction, error) {
	card, err := getCardById(cardId)
	if err != nil {
		return model.Transaction{}, err
	}

	return Withdraw(card.AccountID, CardAccountId, FeeAccountId, amount, model.TransactionTypeVirtualCardDeposit,
		orderId, &cardId)
}

func WithdrawCard(cardId uint, amount uint, orderId string) (model.Transaction, error) {
	card, err := getCardById(cardId)
	if err != nil {
		return model.Transaction{}, err
	}

	return Deposit(CardAccountId, card.AccountID, FeeAccountId, amount, model.TransactionTypeVirtualCardWithdraw,
		orderId, &cardId)
}

func getCardById(id uint) (model.Card, error) {
	card := model.Card{}
	result := model.DB.First(&card, id)
	return card, result.Error
}
