package validator

import (
	"testing"

	"github.com/faparicioaldo/bank-account-practice/validator"
)

func TestAccountNumberValidate(t *testing.T) {
	isValid := validator.AccountNumberValidate("345882865")
	if !isValid {
		t.Errorf("The account number is not valid")
	}

}
