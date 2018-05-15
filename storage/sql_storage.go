package storage

import (
	"github.com/jinzhu/gorm"
)


type sqlStorage struct {
	storage *gorm.DB
}


func NewSQLStorage(sqlPath string) (*sqlStorage, error){
	dbStorage, err := InitDB("sqlite3", sqlPath)
	if err != nil{
		return nil, err
	}

	return &sqlStorage{storage: dbStorage}, nil
}

func (s sqlStorage) Close() error {
	return  s.storage.Close()
}


func (s *sqlStorage) GetGiphs() (Giphs, error){
	var giphs Giphs
	return giphs, s.storage.Find(&giphs).Error
}

