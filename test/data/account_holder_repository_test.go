package bankcardnumber

import (
	"github.com/faparicioaldo/bank-account-practice/data"
	"github.com/faparicioaldo/bank-account-practice/entity"
	"testing"
)

var (
	accountHolder entity.AccountHolder
)

func init() {
	accountHolder = entity.AccountHolder{
		Name:  "Juan Perez",
		Email: "juan@juan.com",
		Alias: "Juanito",
	}
}

func TestGetAccountHoldersCount(t *testing.T) {
	tx := data.DB.MustBegin()
	defer tx.Rollback()
	_, err := data.GetAccountHoldersCount(tx)
	if err != nil {
		t.Errorf("It counldn't get account holders count, error: %v", err)
	}
}

func TestAddAccountHolder(t *testing.T) {
	tx := data.DB.MustBegin()
	defer tx.Rollback()
	id := data.AddAccountHolder(tx, &accountHolder)
	if id == 0 {
		t.Error("User could't be created")
	}
}
