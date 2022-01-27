package validator

import (
	"strconv"
	"strings"
)

func AccountNumberValidate(accountNumber string) bool {
	var isValidAccount bool
	accountSplit := strings.Split(accountNumber, "")
	last := len(accountSplit)
	total := 0
	for i := 0; i < last; i++ {
		item := accountSplit[(last-1)-i]
		itemInt, _ := strconv.Atoi(item)

		di := itemInt * (i + 1)
		total += di
	}
	res := total % 11
	if res == 0 {
		isValidAccount = true
	}

	return isValidAccount
}

func ValidBankID(bankID int) bool {
	switch bankID {
	case 11, 22, 33, 44, 55, 66, 77:
		return true
	default:
		return false
	}
}
