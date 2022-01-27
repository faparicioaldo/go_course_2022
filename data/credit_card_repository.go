package data

import (
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/jmoiron/sqlx"
)

func GetAllCreditCards(tx *sqlx.Tx) (creditCards []*entity.CreditCard, err error) {
	err = tx.Select(&creditCards, "SELECT * FROM credit_cards ORDER BY RANDOM()")
	return
}

func GetCreditCardById(tx *sqlx.Tx, creditCardId int) (creditCard entity.CreditCard, err error) {
	err = tx.Get(&creditCard, "SELECT * FROM credit_cards WHERE id = $1", creditCardId)
	return
}

func AddCreditCard(tx *sqlx.Tx, card *entity.CreditCard) (err error) {
	_, err = tx.NamedExec("INSERT INTO credit_cards (card_number, id_account_holder, id_bank_code, id_product_segment, id_expedition_country) VALUES (:card_number, :id_account_holder, :id_bank_code, :id_product_segment, :id_expedition_country)", card)
	return
}

func DeleteCreditCard(tx *sqlx.Tx, creditCardNumber uint64) (count int64, err error) {
	res, err := tx.Exec("DELETE FROM credit_cards WHERE id = $1", creditCardNumber)
	if err == nil {
		count, err = res.RowsAffected()
	}
	return
}

func UpdateCreditCard(tx *sqlx.Tx, creditCard *entity.CreditCard) (err error) {
	_, err = tx.NamedExec("UPDATE credit_cards SET card_number = :card_number, id_account_holder = :id_account_holder, id_bank_code = :id_bank_code, id_product_segment = :id_product_segment, id_expedition_country = :id_expedition_country WHERE id = :id", creditCard)
	return
}
