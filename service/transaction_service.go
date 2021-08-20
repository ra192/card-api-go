package service

import (
	"errors"
	"github.com/ra192/card-api-go/model"
	"log"
	"time"
)

func Deposit(srcAccountId uint, destAccountId uint, feeAccountId uint, amount uint, transactionType string, orderId string, cardId *uint) (model.Transaction, error) {
	srcAccount, err := GetActiveAccountById(srcAccountId)
	if err != nil {
		return model.Transaction{}, err
	}
	destAccount, err := GetActiveAccountById(destAccountId)
	if err != nil {
		return model.Transaction{}, err
	}

	feeAmount := calculateFee(amount, transactionType, destAccountId)
	if sumByAccount(srcAccountId)-amount < 0 {
		return model.Transaction{}, errors.New("source account does not have enough funds")
	}

	transaction, err := createTransaction(srcAccount, destAccount, amount, transactionType, orderId, cardId)
	if err != nil {
		return model.Transaction{}, err
	}

	if feeAmount > 0 {
		err = model.DB.Create(&model.TransactionItem{
			Amount:               feeAmount,
			Created:              time.Time{},
			TransactionID:        transaction.ID,
			SourceAccountID:      destAccountId,
			DestinationAccountID: feeAccountId,
			CardID:               nil,
		}).Error

		if err != nil {
			return model.Transaction{}, err
		}
	}

	log.Printf("Transaction %s was created", transaction.Type)

	return transaction, nil
}

func Fund(accountId uint, amount uint, orderId string) (model.Transaction, error) {
	cashAccount, err := GetCashAccount()
	if err != nil {
		return model.Transaction{}, err
	}

	account, err := GetActiveAccountById(accountId)
	if err != nil {
		return model.Transaction{}, err
	}

	transaction, err := createTransaction(cashAccount, account, amount, model.TransactionTypeFund, orderId, nil)

	if err != nil {
		return model.Transaction{}, err
	}

	log.Printf("Transaction %s was created", transaction.Type)

	return transaction, nil
}

func Withdraw(srcAccountId uint, destAccountId uint, feeAccountId uint, amount uint, transactionType string, orderId string, cardId *uint) (model.Transaction, error) {
	srcAccount, err := GetActiveAccountById(srcAccountId)
	if err != nil {
		return model.Transaction{}, err
	}
	destAccount, err := GetActiveAccountById(destAccountId)
	if err != nil {
		return model.Transaction{}, err
	}

	feeAmount := calculateFee(amount, transactionType, srcAccountId)
	if sumByAccount(srcAccountId)-amount-feeAmount < 0 {
		return model.Transaction{}, errors.New("source account does not have enough funds")
	}

	transaction, err := createTransaction(srcAccount, destAccount, amount, transactionType, orderId, cardId)
	if err != nil {
		return model.Transaction{}, err
	}

	if feeAmount > 0 {
		err = model.DB.Create(&model.TransactionItem{
			Amount:               feeAmount,
			Created:              time.Time{},
			TransactionID:        transaction.ID,
			SourceAccountID:      srcAccountId,
			DestinationAccountID: feeAccountId,
			CardID:               nil,
		}).Error
		if err != nil {
			return model.Transaction{}, err
		}
	}

	log.Printf("Transaction %s was created", transaction.Type)

	return transaction, nil
}

func createTransaction(srcAccount model.Account, destAccount model.Account, amount uint, transactionType string,
	orderId string, cardId *uint) (model.Transaction, error) {
	if srcAccount.Currency != destAccount.Currency {
		return model.Transaction{}, errors.New("source account currency doesn't match destination account currency")
	}

	Transaction := model.Transaction{
		OrderId: orderId,
		Type:    transactionType,
		Status:  model.TransactionStatusCompleted,
	}

	model.DB.Create(&Transaction)

	model.DB.Create(&model.TransactionItem{
		Amount:               amount,
		Created:              time.Now(),
		TransactionID:        Transaction.ID,
		SourceAccountID:      srcAccount.ID,
		DestinationAccountID: destAccount.ID,
		CardID:               cardId,
	})

	return Transaction, nil
}

func calculateFee(amount uint, transactionType string, accountId uint) uint {
	fee := model.TransactionFee{}
	result := model.DB.First(&fee, model.TransactionFee{Type: transactionType, AccountID: accountId})
	if result.Error != nil {
		return 0
	}

	return uint(float32(amount) * fee.Rate)
}

func sumByAccount(accountId uint) uint {
	var destAmount uint
	res := model.DB.Model(&model.TransactionItem{}).Select("sum(amount)").
		Where(&model.TransactionItem{DestinationAccountID: accountId}).Scan(&destAmount)

	if res.Error != nil {
		destAmount = 0
	}

	var srcAmount uint
	res = model.DB.Model(&model.TransactionItem{}).Select("sum(amount)").
		Where(&model.TransactionItem{SourceAccountID: accountId}).Scan(&srcAmount)

	if res.Error != nil {
		srcAmount = 0
	}

	return destAmount - srcAmount
}
