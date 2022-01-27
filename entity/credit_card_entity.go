package entity

type CreditCard struct {
	Id                  int    `db:"id"`
	CardNumber          string `db:"card_number"`
	AccountHolderId     int    `db:"id_account_holder"`
	BankCodeId          int    `db:"id_bank_code"`
	SegmentId           int    `db:"id_product_segment"`
	ExpeditionCountryId int    `db:"id_expedition_country"`
}
