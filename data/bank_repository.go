package data

import (
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/jmoiron/sqlx"
)

func GetAllBanks(tx *sqlx.Tx) (banks []*entity.Bank, err error) {
	err = tx.Select(&banks, "SELECT * FROM bank_codes ORDER BY RANDOM()")
	return
}

func GetBankById(tx *sqlx.Tx, bankId int) (bank entity.Bank, err error) {
	err = tx.Get(&bank, "SELECT * FROM bank_codes WHERE id = $1", bankId)
	return
}

func GetRandomBank(tx *sqlx.Tx) (bank entity.Bank, err error) {
	err = tx.QueryRowx("SELECT * FROM bank_codes ORDER BY RANDOM() LIMIT 1").StructScan(&bank)
	return
}

func AddBank(tx *sqlx.Tx, bank *entity.Bank) (err error) {
	_, err = tx.NamedExec("INSERT INTO bank_codes (code, name) VALUES (:code, :name)", bank)
	return
}

func DeleteBank(tx *sqlx.Tx, bankId int) (count int64, err error) {
	res, err := tx.Exec("DELETE FROM bank_codes WHERE id = $1", bankId)
	if err == nil {
		count, err = res.RowsAffected()
	}
	return
}

func UpdateBank(tx *sqlx.Tx, bank *entity.Bank) (err error) {
	_, err = tx.NamedExec("UPDATE bank_codes SET code = :code, name = :name WHERE id = :id", bank)
	return
}
