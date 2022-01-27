package api

import (
	"github.com/faparicioaldo/bank-account-practice/handler"
	"github.com/gin-gonic/gin"
)

func CreateCreditCard(g *gin.Engine) {
	g.GET("/api/v1/credit_cards/:id", handler.GetCreditCardById)
	g.GET("/api/v1/credit_cards", handler.GetCreditCards)
	g.POST("/api/v1/credit_cards", handler.AddCreditCard)
	g.DELETE("/api/v1/credit_cards/:id", handler.DeleteCreditCard)
	g.PUT("/api/v1/credit_cards/:id", handler.UpdateCreditCard)
}
