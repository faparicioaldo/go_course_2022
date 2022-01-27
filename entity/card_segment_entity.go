package entity

type ProductSegment struct {
	Id   int    `db:"id"`
	Code int    `db:"code"`
	Name string `db:"name"`
}
