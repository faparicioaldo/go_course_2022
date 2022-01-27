package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/gin-gonic/gin"
)

func GetAccountHoldersCount(ctx *gin.Context) {
	tx := data.DB.MustBegin()
	count, err := data.GetAccountHoldersCount(tx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "A problem ocurrs to get account holders count, try later",
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"accountHoldersCount": count,
	})
}

func GetAccountHolderById(ctx *gin.Context) {
	accountHolderId := ctx.Param("id")
	accountHolderId_Int, _ := strconv.Atoi(accountHolderId)
	tx := data.DB.MustBegin()
	accountHolder, err := data.GetAccountHolderById(tx, accountHolderId_Int)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"accountHolder": accountHolder,
	})
}

func GetAccountHolders(ctx *gin.Context) {
	tx := data.DB.MustBegin()
	accountHolders, err := data.GetAllAccountHolders(tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ping Pong",
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"accountHolders": accountHolders,
	})
}

func AddAccountHolder(ctx *gin.Context) {
	accountHolder := &entity.AccountHolder{}
	if err := ctx.ShouldBindJSON(accountHolder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := data.DB.MustBegin()
	var id int
	if id, _ = data.AddAccountHolder(tx, accountHolder); id == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("Can't insert account holder %v", id),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": fmt.Sprintf("AccountHolder created successfully  with ID %v", id),
	})
}

func DeleteAccountHolder(ctx *gin.Context) {
	accountHolderId := ctx.Param("id")
	accountHolderId_Int, _ := strconv.Atoi(accountHolderId)
	tx := data.DB.MustBegin()
	var rowsAffected int64
	var err error
	if rowsAffected, err = data.DeleteAccountHolder(tx, accountHolderId_Int); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": fmt.Sprintf("AccountHolder deleted successfully %v", rowsAffected),
	})
}

func UpdateAccountHolder(ctx *gin.Context) {
	accountHolderId := ctx.Param("id")
	accountHolderId_Int, _ := strconv.Atoi(accountHolderId)

	accountHolder := &entity.AccountHolder{}
	if err := ctx.ShouldBindJSON(accountHolder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := data.DB.MustBegin()
	accountHolder.Id = accountHolderId_Int

	if err := data.UpdateAccountHolder(tx, accountHolder); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "AccountHolder was updated sucessfully",
	})
}

func HolaAldo(ctx *gin.Context) {
	log.Println("Hola Aldo")
}

func AdiosAldo(ctx *gin.Context) {
	log.Println("Adios Aldo")
}
