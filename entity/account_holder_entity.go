package entity

type AccountHolder struct {
	Id    int    `db:"id"`
	Name  string `db:"name"`
	Email string `db:"email"`
	Alias string `db:"alias"`
}
