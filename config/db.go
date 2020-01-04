package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var screatkey = []byte("bian")

//type PostgresDBActivity struct {
//	DB *gorm.DB
//}

func DBInit() *gorm.DB {

	//db, err := gorm.Open("mysql", "nanda:nanda@tcp(192.168.133.87:3306)/warnadev?charset=utf8&parseTime=True&loc=Local")
	//db, err := gorm.Open("postgres", "host=dna-developer.com port=5432 user=postgres dbname=warnadev password=nanda")
	db, err := gorm.Open("postgres", "host=192.168.133.89 port=5432 user=postgres dbname=warnadev password=nanda")
	//db, err := gorm.Open("postgres", "host=192.168.133.81 port=5432 user=postgres dbname=warnaprod password=nandaprod")
	if err != nil {

		// panic("Gagal koneksi ke database")
		panic(err.Error())
		//return err
		//return err.Error()
	}
	//dbS.DB = db

	return db
}

func JwtKey() []byte {

	return screatkey
}
