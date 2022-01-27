package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/gin-gonic/gin"
)

func GetBankById(ctx *gin.Context) {
	bankId := ctx.Param("id")
	bankId_Int, _ := strconv.Atoi(bankId)
	tx := data.DB.MustBegin()
	bank, err := data.GetBankById(tx, bankId_Int)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"bank": bank,
	})
}

func GetBanks(ctx *gin.Context) {
	tx := data.DB.MustBegin()
	banks, err := data.GetAllBanks(tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ping Pong",
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"banks": banks,
	})
}

func AddBank(ctx *gin.Context) {
	bank := &entity.Bank{}
	if err := ctx.ShouldBindJSON(bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := data.DB.MustBegin()
	if err := data.AddBank(tx, bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "Bank created successfully",
	})
}

func DeleteBank(ctx *gin.Context) {
	bankId := ctx.Param("id")
	bankId_Int, _ := strconv.Atoi(bankId)
	tx := data.DB.MustBegin()
	var rowsAffected int64
	var err error
	if rowsAffected, err = data.DeleteBank(tx, bankId_Int); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": fmt.Sprintf("Bank deleted successfully %v", rowsAffected),
	})
}

func UpdateBank(ctx *gin.Context) {
	bankId := ctx.Param("id")
	bankId_Int, _ := strconv.Atoi(bankId)

	bank := &entity.Bank{}
	if err := ctx.ShouldBindJSON(bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := data.DB.MustBegin()
	bank.Id = bankId_Int

	if err := data.UpdateBank(tx, bank); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Bank was updated sucessfully",
	})
}
