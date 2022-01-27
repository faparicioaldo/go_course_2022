package data

import (
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/jmoiron/sqlx"
)

func GetAllExpeditionCountries(tx *sqlx.Tx) (expeditionCountries []*entity.ExpeditionCountry, err error) {
	err = tx.Select(&expeditionCountries, "SELECT * FROM expedition_countries ORDER BY RANDOM()")
	return
}

func GetRandomExpeditionCountry(tx *sqlx.Tx) (country entity.ExpeditionCountry, err error) {
	err = tx.QueryRowx("SELECT * FROM expedition_countries ORDER BY RANDOM() LIMIT 1").StructScan(&country)
	return
}

func GetExpeditionCountryById(tx *sqlx.Tx, bankId int) (expeditionCountry entity.ExpeditionCountry, err error) {
	err = tx.Get(&expeditionCountry, "SELECT * FROM expedition_countries WHERE id = $1", bankId)
	return
}

func AddExpeditionCountry(tx *sqlx.Tx, expeditionCountry *entity.ExpeditionCountry) (err error) {
	_, err = tx.NamedExec("INSERT INTO expedition_countries (code, name) VALUES (:code, :name)", expeditionCountry)
	return
}

func DeleteExpeditionCountry(tx *sqlx.Tx, bankId int) (count int64, err error) {
	res, err := tx.Exec("DELETE FROM expedition_countries WHERE id = $1", bankId)
	if err == nil {
		count, err = res.RowsAffected()
	}
	return
}

func UpdateExpeditionCountry(tx *sqlx.Tx, expeditionCountry *entity.ExpeditionCountry) (err error) {
	_, err = tx.NamedExec("UPDATE expedition_countries SET code = :code, name = :name WHERE id = :id", expeditionCountry)
	return
}
