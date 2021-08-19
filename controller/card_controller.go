package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ra192/card-api-go/model"
	"github.com/ra192/card-api-go/service"
	"net/http"
	"time"
)

type CreateCardDto struct {
	AccountId uint `json:"accountId"`
	CustomerId uint `json:"customerId"`
}

func CreateVirtualCard(c *gin.Context) {
	_, err := validateToken(c)
	if err!=nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req CreateCardDto
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	card := model.Card{AccountID: req.AccountId, CustomerID: req.CustomerId, Created: time.Now(), Type: model.CardTypeVirtual}
	service.СreateCard(&card)

	c.JSON(http.StatusOK, gin.H{"cardId": card.ID})
}
