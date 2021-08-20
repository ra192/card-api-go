package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ra192/card-api-go/model"
	"github.com/ra192/card-api-go/service"
	"net/http"
	"time"
)

type CreateCardDto struct {
	AccountId  uint `json:"accountId"`
	CustomerId uint `json:"customerId"`
}

func CreateVirtualCard(c *gin.Context) {
	_, err := validateToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req CreateCardDto
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	card := model.Card{AccountID: req.AccountId, CustomerID: req.CustomerId, Created: time.Now(), Type: model.CardTypeVirtual}
	err = service.СreateCard(&card)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"cardId": card.ID})
}

type CardTransactionDto struct {
	CardId  uint   `json:"cardId"`
	Amount  uint   `json:"amount"`
	OrderId string `json:"orderId"`
}

func DepositCard(c *gin.Context) {
	_, err := validateToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req CardTransactionDto
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := service.DepositCard(req.CardId, req.Amount, req.OrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionId": transaction.ID})
}

func WithdrawCard(c *gin.Context) {
	_, err := validateToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req CardTransactionDto
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	transaction, err := service.WithdrawCard(req.CardId, req.Amount, req.OrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionId": transaction.ID})
}
