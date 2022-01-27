package handler

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/faparicioaldo/bank-account-practice/generator"
	"github.com/faparicioaldo/bank-account-practice/validator"
	"github.com/gin-gonic/gin"
)

const (
	csvFilePath = "/Users/javadabadoo/repositorio/go/src/gitlab.com/java.daba.doo/go-practice/test/resources/users.csv"
)

func GetCreditCardById(ctx *gin.Context) {
	creditCardId := ctx.Param("id")
	creditCardId_Int, _ := strconv.Atoi(creditCardId)
	tx := data.DB.MustBegin()
	creditCard, err := data.GetCreditCardById(tx, creditCardId_Int)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"credit_card": creditCard,
	})
}

func GetCreditCards(ctx *gin.Context) {
	tx := data.DB.MustBegin()
	creditCards, err := data.GetAllCreditCards(tx)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "Ping Pong",
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"credit_cards": creditCards,
	})
}

func AddCreditCard(ctx *gin.Context) {
	creditCard := &entity.CreditCard{}
	if err := ctx.ShouldBindJSON(creditCard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx := data.DB.MustBegin()
	if err := data.AddCreditCard(tx, creditCard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": "Credit card created successfully",
	})
}

func DeleteCreditCard(ctx *gin.Context) {
	creditCardNumber := ctx.Param("card_number")
	creditCardNumber_Int, _ := strconv.ParseUint(creditCardNumber, 10, 64)
	tx := data.DB.MustBegin()
	var rowsAffected int64
	var err error
	if rowsAffected, err = data.DeleteCreditCard(tx, creditCardNumber_Int); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusCreated, gin.H{
		"status": fmt.Sprintf("Credit card deleted successfully %v", rowsAffected),
	})
}

func UpdateCreditCard(ctx *gin.Context) {
	creditCardNumber := ctx.Param("card_number")
	// creditCardNumber_Int, _ := strconv.ParseUint(creditCardNumber, 10, 64)

	creditCard := &entity.CreditCard{}
	if err := ctx.ShouldBindJSON(creditCard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	tx := data.DB.MustBegin()
	creditCard.CardNumber = creditCardNumber

	if err := data.UpdateCreditCard(tx, creditCard); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, gin.H{
		"status": "Credit card was updated sucessfully",
	})
}

func CreateCreditCardFromAccountHolder(user entity.AccountHolder) {
	log.Println("CreateCard: init")
	defer log.Println("CreateCard: end")
	tx := data.DB.MustBegin()
	randomBank, _ := data.GetRandomBank(tx)
	randomCardSegment, _ := data.GetRandomProductSegment(tx)
	randomCountry, _ := data.GetRandomExpeditionCountry(tx)

	cardNumberPrefix := strconv.Itoa(randomBank.Code) +
		strconv.Itoa(randomCardSegment.Code) +
		strconv.Itoa(randomCountry.Code)
	var randomAccountNumber string

	for {
		randomAccountNumber = generator.GenerateRandomAccountNumber()
		if validator.AccountNumberValidate(randomAccountNumber) {
			log.Printf("\nNumber %v : Valid with lenght: %v", randomAccountNumber, len(randomAccountNumber))
			break
		}
	}

	newCardNumber := cardNumberPrefix + randomAccountNumber

	log.Println("randomBank: ", randomBank)
	log.Println("randomCardSegment: ", randomCardSegment)
	log.Println("randomCountry: ", randomCountry)
	log.Println("newCardNumber: ", newCardNumber)

	id, _ := data.AddAccountHolder(tx, &user)

	newCreditCard := entity.CreditCard{
		CardNumber:          newCardNumber,
		AccountHolderId:     id,
		BankCodeId:          randomBank.Id,
		SegmentId:           randomCardSegment.Id,
		ExpeditionCountryId: randomCountry.Id,
	}
	err2 := data.AddCreditCard(tx, &newCreditCard)

	log.Println("newCreditCard: ", newCreditCard)

	if id == 0 {
		log.Println("ERROR", id)
	}

	if err2 != nil {
		log.Println("ERROR 2", err2)
	}
	tx.Commit()
}
