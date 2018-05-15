package storage

import (
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(dbtype string, dbname string) (*gorm.DB, error) {
	db, err := gorm.Open(dbtype, dbname)
	if err != nil{
		log.Print(err)
		return nil, err
	}

	db.LogMode(true)

	if !db.HasTable(&Giph{}){
		if err:= db.CreateTable(&Giph{}).Set("gorm:table_options", "ENGINE=InnoDB").Error; err != nil {
			return nil, err
		}
	}

	return db, nil
}

