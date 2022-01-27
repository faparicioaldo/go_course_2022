package api

import (
	"github.com/faparicioaldo/bank-account-practice/handler"
	"github.com/gin-gonic/gin"
)

func CreateExpeditionCountry(g *gin.Engine) {
	g.GET("/api/v1/expedition_countries/:id", handler.HolaAldo, handler.GetExpeditionCountryById, handler.AdiosAldo)
	// g.GET("/api/v1/expedition_countries/:id", handler.GetExpeditionCountryById)
	g.GET("/api/v1/expedition_countries", handler.GetExpeditionCountries)
	g.POST("/api/v1/expedition_countries", handler.AddExpeditionCountry)
	g.DELETE("/api/v1/expedition_countries/:id", handler.DeleteExpeditionCountry)
	g.PUT("/api/v1/expedition_countries/:id", handler.UpdateExpeditionCountry)

}
