package api

import (
	"github.com/faparicioaldo/bank-account-practice/handler"
	"github.com/gin-gonic/gin"
)

func CreateBank(g *gin.Engine) {
	g.GET("/api/v1/banks/:id", handler.GetBankById)
	g.GET("/api/v1/banks", handler.GetBanks)
	g.POST("/api/v1/banks", handler.AddBank)
	g.DELETE("/api/v1/banks/:id", handler.DeleteBank)
	g.PUT("/api/v1/banks/:id", handler.UpdateBank)
}
