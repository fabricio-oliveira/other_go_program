package user

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func InitDB(uri string) (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", uri)
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	db.AutoMigrate(&Model{})
	db.DB().SetMaxOpenConns(5)
	return db, nil
}
