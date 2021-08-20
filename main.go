package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ra192/card-api-go/controller"
	"github.com/ra192/card-api-go/model"
)

func main() {
	r := gin.Default()

	r.POST("/api/token", controller.CreateToken)
	r.POST("/api/account/fund", controller.FundAccount)
	r.POST("/api/customer", controller.CreateCustomer)
	r.POST("/api/card", controller.CreateVirtualCard)
	r.POST("/api/card/deposit", controller.DepositCard)
	r.POST("/api/card/withdraw", controller.WithdrawCard)

	model.ConnectDataBase()

	r.Run()
}
