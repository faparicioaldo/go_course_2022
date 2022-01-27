package data

import (
	"log"

	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/jmoiron/sqlx"
)

func GetAccountHoldersCount(tx *sqlx.Tx) (count int, err error) {
	row := tx.QueryRow("SELECT count(*) FROM account_holders")
	err = row.Scan(&count)
	if err != nil {
		log.Println("Couldn't get account holders count")
	}
	return
}

func GetAllAccountHolders(tx *sqlx.Tx) (accountHolders []*entity.AccountHolder, err error) {
	err = tx.Select(&accountHolders, "SELECT * FROM account_holders ORDER BY RANDOM()")
	return
}

func GetRandomAccountHolder() (tx *sqlx.Tx, accountHolder entity.AccountHolder, err error) {
	err = tx.QueryRowx("SELECT * FROM account_holders ORDER BY RANDOM() LIMIT 1").StructScan(&accountHolder)
	return
}

func GetAccountHolderById(tx *sqlx.Tx, accountHolderId int) (accountHolder entity.AccountHolder, err error) {
	err = tx.Get(&accountHolder, "SELECT * FROM account_holders WHERE id = $1", accountHolderId)
	return
}

func AddAccountHolder(tx *sqlx.Tx, accountHolder *entity.AccountHolder) (int, error) {
	res, err := tx.PrepareNamed("INSERT INTO account_holders (name, email, alias) VALUES (:name, :email, :alias) RETURNING id")
	if err != nil {
		log.Println(err)
	}
	var id int
	err = res.Get(&id, accountHolder)
	if err != nil {
		log.Println(err)
	}
	return id, err
}

func DeleteAccountHolder(tx *sqlx.Tx, accountHolderId int) (count int64, err error) {
	res, err := tx.Exec("DELETE FROM account_holders WHERE id = $1", accountHolderId)
	if err == nil {
		count, err = res.RowsAffected()
	}
	return
}

func UpdateAccountHolder(tx *sqlx.Tx, accountHolder *entity.AccountHolder) (err error) {
	_, err = tx.NamedExec("UPDATE account_holders SET name = :name, email = :email, alias = :alias WHERE id = :id", accountHolder)
	return
}
