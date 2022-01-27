package data

import (
	"github.com/faparicioaldo/bank-account-practice/entity"
	"github.com/jmoiron/sqlx"
)

func GetAllProductSegments(tx *sqlx.Tx) (productSegments []*entity.ProductSegment, err error) {
	err = tx.Select(&productSegments, "SELECT * FROM product_segments ORDER BY RANDOM()")
	return
}

func GetRandomProductSegment(tx *sqlx.Tx) (productSegment entity.ProductSegment, err error) {
	err = tx.QueryRowx("SELECT * FROM product_segments ORDER BY RANDOM() LIMIT 1").StructScan(&productSegment)
	return
}

func GetProductSegmentById(tx *sqlx.Tx, productSegmentId int) (productSegment entity.ProductSegment, err error) {
	err = tx.Get(&productSegment, "SELECT * FROM product_segments WHERE id = $1", productSegmentId)
	return
}

func AddProductSegment(tx *sqlx.Tx, productSegment *entity.ProductSegment) (err error) {
	_, err = tx.NamedExec("INSERT INTO product_segments (code, name) VALUES (:code, :name)", productSegment)
	return
}

func DeleteProductSegment(tx *sqlx.Tx, productSegmentId int) (count int64, err error) {
	res, err := tx.Exec("DELETE FROM product_segments WHERE id = $1", productSegmentId)
	if err == nil {
		count, err = res.RowsAffected()
	}
	return
}

func UpdateProductSegment(tx *sqlx.Tx, productSegment *entity.ProductSegment) (err error) {
	_, err = tx.NamedExec("UPDATE product_segments SET code = :code, name = :name WHERE id = :id", productSegment)
	return
}
