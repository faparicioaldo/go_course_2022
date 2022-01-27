package entity

import "fmt"

type Bank struct {
	Id   int    `db:"id"`
	Code int    `db:"code"`
	Name string `db:"name"`
}

func (b *Bank) String() string {
	return fmt.Sprintf("Bank#String{Id: %d, Code: %d, Name: %s}", b.Id, b.Code, b.Name)
}
