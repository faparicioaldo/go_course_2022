package main

import (
	"log"

	"github.com/faparicioaldo/bank-account-practice/api"
	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/handler"
	"github.com/faparicioaldo/bank-account-practice/reader"
	"github.com/gin-gonic/gin"
)

const (
	csvFilePath = "./test/resources/users.csv"
)

var (
	g *gin.Engine
)

func init() {
	g = gin.Default()
	if getAccountHoldersCount() == 0 {
		loadUsers()
	}
}

func main() {
	api.CreateBank(g)
	api.CreateCreditCard(g)
	api.CreateProductSegment(g)
	api.CreateExpeditionCountry(g)
	api.CreateAccountHolder(g)
	g.Run()
}

func getAccountHoldersCount() int {
	tx := data.DB.MustBegin()
	defer tx.Rollback()
	count, err := data.GetAccountHoldersCount(tx)
	if err != nil {
		log.Printf("It counldn't get account holders count, error: %v", err)
	}
	return count
}

func loadUsers() {
	var reader reader.Reader = &reader.CSVReader{}
	users, _ := reader.ReadFromFile(csvFilePath)

	for _, user := range users {
		handler.CreateCreditCardFromAccountHolder(user)
	}
}
