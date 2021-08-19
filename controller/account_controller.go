package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ra192/card-api-go/service"
	"net/http"
)

const InternalMerchantId = 1

type FundAccountDto struct {
	AccountId uint   `json:"accountId"`
	Amount    uint   `json:"amount"`
	OrderId   string `json:"orderId"`
}

func FundAccount(c *gin.Context) {
	merchantId, err := validateToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if merchantId != InternalMerchantId {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal merchant required"})
		return
	}

	var req FundAccountDto
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	trans, err := service.Fund(req.AccountId, req.Amount, req.OrderId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"transactionId": trans.ID})
}
