package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ra192/card-api-go/service"
	"net/http"
	"strings"
)

type CreateTokenDto struct {
	MerchantId uint   `json:"merchantId"`
	Secret     string `json:"secret"`
}

func CreateToken(c *gin.Context) {
	var req CreateTokenDto

	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := service.CreateToken(req.MerchantId, req.Secret)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}

func validateToken(c *gin.Context) (uint, error) {
	token := strings.Trim(strings.Replace(c.GetHeader("Authorization"), "Bearer", "", 1), " ")
	return service.ValidateToken(token)
}
