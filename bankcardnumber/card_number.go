package bankcardnumber

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/faparicioaldo/bank-account-practice/entity"
)

var banks map[string]string = map[string]string{
	"11": "Banca Bunsan",
	"22": "Go Bank",
	"33": "Aldo Institución Financiera",
	"44": "Angel Banca Privada",
	"55": "Hugo Finantial Services",
	"66": "BanHumberto",
	"77": "Gustavo Commercial Bank",
}

var cardSegments map[string]string = map[string]string{
	"2": "Tarjeta básica",
	"4": "Tarjeta departamental",
	"6": "Tarjeta clásica",
	"8": "Tarjeta Oro",
	"0": "Tarjeta Platino",
}

var countryMx string = "5211"

func GeneratePrefix() string {
	/*
	 */
	return entity.GeneratePrefix()
}

func GetCard() {
	var accounts []AccountHolder = []AccountHolder{
		{
			Name:  "Gerardo Aquino",
			Email: "gerardo@bunsan.io",
			Alias: "javadabadoo",
			CreditCard: CreditCard{
				CardNumber:  22452000000000,
				SegmentId:   4,
				SegmentName: "",
				BankName:    "Go Bank",
				BankID:      22,
			},
		},
		{
			Name:  "Pedro Picapiedra",
			Email: "pedro@piedradura.stoneage",
			Alias: "peter.rocks",
			CreditCard: CreditCard{
				CardNumber:  11052000000000,
				SegmentId:   0,
				SegmentName: "",
				BankName:    "Banca Bunsan",
				BankID:      11,
			},
		},
	}

	b, err := json.MarshalIndent(accounts, "", "   ")
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
}
