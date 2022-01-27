package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/gin-gonic/gin"
)

func GetExpeditionCountryById(ctx *gin.Context) {
	expeditionCountryId := ctx.Param("id")
	expeditionCountryId_Int, _ := strconv.Atoi(expeditionCountryId)
	tx := data.DB.MustBegin()
	expeditionCountry, err := data.GetExpeditionCountryById(tx, expeditionCountryId_Int)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"expedition_country": expeditionCountry,
	})
}

func GetExpeditionCountries(ctx *gin.Context) {
	tx := data.DB.MustBegin()
	expeditionCountrys, err := data.GetAllExpeditionCountries(tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ping Pong",
		})
		return
	}
	tx.Commit()

	ctx.JSON(http.StatusOK, gin.H{
		"expedition_countries": expeditionCountrys,
	})
}

func AddExpeditionCountry(ctx *gin.Context) {
	expeditionCountry := &entity.ExpeditionCountry{}
	if err := ctx.ShouldBindJSON(expeditionCountry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := data.DB.MustBegin()
	if err := data.AddExpeditionCountry(tx, expeditionCountry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "Expedition country created successfully",
	})
}

func DeleteExpeditionCountry(ctx *gin.Context) {
	expeditionCountryId := ctx.Param("id")
	expeditionCountryId_Int, _ := strconv.Atoi(expeditionCountryId)
	tx := data.DB.MustBegin()
	var rowsAffected int64
	var err error
	if rowsAffected, err = data.DeleteExpeditionCountry(tx, expeditionCountryId_Int); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": fmt.Sprintf("Expedition country deleted successfully %v", rowsAffected),
	})
}

func UpdateExpeditionCountry(ctx *gin.Context) {
	expeditionCountryId := ctx.Param("id")
	expeditionCountryId_Int, _ := strconv.Atoi(expeditionCountryId)

	expeditionCountry := &entity.ExpeditionCountry{}
	if err := ctx.ShouldBindJSON(expeditionCountry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := data.DB.MustBegin()
	expeditionCountry.Id = expeditionCountryId_Int

	if err := data.UpdateExpeditionCountry(tx, expeditionCountry); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Expedition country was updated sucessfully",
	})
}
