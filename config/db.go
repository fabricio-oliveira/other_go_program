package config

import (
	"errors"

	"github.com/fabricio-oliveira/other_go_program/user"

	"github.com/jinzhu/gorm"
	//drive database
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

//InitDB init database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open("sqlite3", "./db/foo.db")
	if err != nil {
		return nil, errors.New("failed to connect database")
	}

	db.AutoMigrate(&user.User{})
	db.DB().SetMaxOpenConns(5)
	return db, nil
}
