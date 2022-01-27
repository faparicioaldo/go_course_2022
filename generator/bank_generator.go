package generator

import (
	"log"
	"math/rand"
	"strconv"
	"time"
)

func GenerateCardNumberPrefix(bankID int, cardSegmentID int, countryID int) string {
	return strconv.Itoa(bankID) + strconv.Itoa(cardSegmentID) + strconv.Itoa(countryID)
}

func GenerateRandomAccountNumber() string {
	min := 000_000_000
	max := 999_999_999
	var accountNumberRandom string
	accountNumberRandom += strconv.Itoa(generateRandomIntByRange(min, max))
	log.Println("Account number generated: ", accountNumberRandom)
	return accountNumberRandom
}

func generateRandomIntByRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min+1) + min
}
