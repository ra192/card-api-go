package service

import (
	"errors"
	"github.com/ra192/card-api-go/model"
	"time"
)

func Fund(accountId uint, amount uint, orderId string) (model.Transaction, error) {
	cashAccount, err := GetCashAccount()
	if err != nil {
		return model.Transaction{}, err
	}

	account, err := GetActiveAccountById(accountId)
	if err != nil {
		return model.Transaction{}, err
	}

	return createTransaction(cashAccount, account, amount, model.TransactionTypeFund, orderId, model.Card{})
}

func createTransaction(srcAccount model.Account, destAccount model.Account, amount uint, transactionType string,
	orderId string, card model.Card) (model.Transaction, error) {
	if srcAccount.Currency != destAccount.Currency {
		return model.Transaction{}, errors.New("source account currency doesn't match destination account currency")
	}

	transaction := model.Transaction{
		OrderId: orderId,
		Type:    transactionType,
		Status:  model.TransactionStatusCompleted,
	}

	model.DB.Create(transaction)

	model.DB.Create(model.TransactionItem{
		Amount:               amount,
		Created:              time.Now(),
		TransactionID:        transaction.ID,
		SourceAccountID:      srcAccount.ID,
		DestinationAccountID: destAccount.ID,
		CardID:               card.ID,
	})

	return transaction, nil
}
