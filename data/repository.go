package data

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var (
	DB *sqlx.DB
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "go_practice"
	sslmode  = "disable"
)

func init() {
	log.Println("Initializing database...")
	connData := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, user, password, dbname, sslmode)
	var err error
	DB, err = sqlx.Connect("postgres", connData)
	if err != nil {
		log.Fatal("Couldn't connect to DB", err)
	}
	DB.SetMaxIdleConns(5)
	DB.SetMaxOpenConns(10)
}
