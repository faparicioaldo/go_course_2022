package bankcardnumber

import (
	"testing"

	bank "github.com/faparicioaldo/bank-account-practice/bankcardnumber"
)

func TestBankIsValid(t *testing.T) {
	bankIsValid := bank.ValidBankID(11)
	if !bankIsValid {
		t.Error("This bank is not valid")
	}
}

func TestBankIsNotValid(t *testing.T) {
	bankIsNotValid := bank.ValidBankID(12)
	if bankIsNotValid {
		t.Error("This bank is not valid")
	}
}

func TestAccountValid(t *testing.T) {
	accountIsValid := bank.ValidateAccountNumber("123456789")
	if !accountIsValid {
		t.Error("This account number is not valid")
	}
}

func TestAccountNotValid(t *testing.T) {
	accountIsValid := bank.ValidateAccountNumber("123456788")
	if accountIsValid {
		t.Error("This account number should be not valid")
	}
}

func TestGenerateRandomAccount(t *testing.T) {
	accountNumberRandom1 := bank.GenerateRandomAccountNumber()
	t.Log(accountNumberRandom1)
	if true {
		//	if accountNumberRandom1 != "213" {
		t.Error("No chavo")
	}
}

func TestGeneratePrefix(t *testing.T) {
	accountNumberRandom1 := bank.GeneratePrefix()
	t.Log(accountNumberRandom1)
	if true {
		//	if accountNumberRandom1 != "213" {
		t.Error("No chavo")
	}
}
