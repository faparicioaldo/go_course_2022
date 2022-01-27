package api

import (
	"github.com/faparicioaldo/bank-account-practice/handler"
	"github.com/gin-gonic/gin"
)

func CreateAccountHolder(g *gin.Engine) {
	g.GET("/api/v1/account_holders/:id", handler.GetAccountHolderById)
	g.GET("/api/v1/account_holders", handler.GetAccountHolders)
	g.POST("/api/v1/account_holders", handler.AddAccountHolder)
	g.DELETE("/api/v1/account_holders/:id", handler.DeleteAccountHolder)
	g.PUT("/api/v1/account_holders/:id", handler.UpdateAccountHolder)

}
