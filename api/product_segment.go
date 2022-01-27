package api

import (
	"github.com/faparicioaldo/bank-account-practice/handler"
	"github.com/gin-gonic/gin"
)

func CreateProductSegment(g *gin.Engine) {
	g.GET("/api/v1/product_segments/:id", handler.GetProductSegmentById)
	g.GET("/api/v1/product_segments", handler.GetProductSegments)
	g.POST("/api/v1/product_segments", handler.AddProductSegment)
	g.DELETE("/api/v1/product_segments/:id", handler.DeleteProductSegment)
	g.PUT("/api/v1/product_segments/:id", handler.UpdateProductSegment)

}
