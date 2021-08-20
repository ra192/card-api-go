package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/ra192/card-api-go/model"
	"github.com/ra192/card-api-go/service"
	"gorm.io/datatypes"
	"net/http"
	"time"
)

type CreateCustomerDto struct {
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	BirthDate   string `json:"birthDate"`
	Address     string `json:"address"`
	Address2    string `json:"address2"`
	City        string `json:"city"`
	StateRegion string `json:"stateRegion"`
	Country     string `json:"country"`
	PostalCode  string `json:"postalCode"`
	MerchantID  uint   `json:"merchantId"`
}

func CreateCustomer(c *gin.Context) {
	merchantId, err := validateToken(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var req CreateCustomerDto
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	customer := model.Customer{Phone: req.Phone, Email: req.Email, Active: true, FirstName: req.FirstName,
		LastName: req.LastName, BirthDate: datatypes.Date(birthDate), Address: req.Address, Address2: req.Address2,
		City: req.City, StateRegion: req.StateRegion, Country: req.Country, PostalCode: req.PostalCode,
		MerchantID: merchantId}
	err = service.CreateCustomer(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"id": customer.ID})
}
