package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/gin-gonic/gin"
)

func GetProductSegmentById(ctx *gin.Context) {
	productSegmentId := ctx.Param("id")
	productSegmentId_Int, _ := strconv.Atoi(productSegmentId)
	tx := data.DB.MustBegin()
	productSegment, err := data.GetProductSegmentById(tx, productSegmentId_Int)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"product_segment": productSegment,
	})
}

func GetProductSegments(ctx *gin.Context) {
	tx := data.DB.MustBegin()
	productSegments, err := data.GetAllProductSegments(tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ping Pong",
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"product_segment": productSegments,
	})
}

func AddProductSegment(ctx *gin.Context) {
	bank := &entity.ProductSegment{}
	if err := ctx.ShouldBindJSON(bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := data.DB.MustBegin()
	if err := data.AddProductSegment(tx, bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "Product segment created successfully",
	})
}

func DeleteProductSegment(ctx *gin.Context) {
	bankId := ctx.Param("id")
	bankId_Int, _ := strconv.Atoi(bankId)
	tx := data.DB.MustBegin()
	var rowsAffected int64
	var err error
	if rowsAffected, err = data.DeleteProductSegment(tx, bankId_Int); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": fmt.Sprintf("Product segment deleted successfully %v", rowsAffected),
	})
}

func UpdateProductSegment(ctx *gin.Context) {
	bankId := ctx.Param("id")
	bankId_Int, _ := strconv.Atoi(bankId)

	bank := &entity.ProductSegment{}
	if err := ctx.ShouldBindJSON(bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := data.DB.MustBegin()
	bank.Id = bankId_Int

	if err := data.UpdateProductSegment(tx, bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Product segment was updated sucessfully",
	})
}
