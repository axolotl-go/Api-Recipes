package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN = "user=postgres.pfjmzstmbubjprmeztxa password=Ny#ZR0g3W8#s98,8{TAP;$C.QBZ host=aws-0-us-west-1.pooler.supabase.com port=5432 dbname=postgres"
var DB *gorm.DB

func Dbconnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
	} else {
		log.Println("DB connection established")
	}
}
