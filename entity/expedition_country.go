package entity

import "fmt"

type ExpeditionCountry struct {
	Id   int    `db:"id"`
	Code int    `db:"code"`
	Name string `db:"name"`
}

func (ec *ExpeditionCountry) String() string {
	return fmt.Sprintf("ExpeditionCountry#String{Id: %d, Code: %d, Name: %s}", ec.Id, ec.Code, ec.Name)
}
