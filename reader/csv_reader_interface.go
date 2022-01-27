package reader

import (
	"github.com/faparicioaldo/bank-account-practice/entity"
)

type Reader interface {
	ReadFromFile(filePath string) (users []entity.AccountHolder, err error)
}
