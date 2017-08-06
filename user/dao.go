package dao

import (
	"log"

	"github.com/jinzhu/gorm"
)

//User UserDAO
type UserDAO struct {
	db *gorm.DB
}

//NewUser create new use
func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{db: db}
}

//Find return a user
func (u UserDAO) Find(id int) (*User, error) {
	user := &User{}
	if erro := u.db.Where("id == ?", id).First(&user).Error; erro != nil {
		return nil, erro
	}
	return user, nil
}

//Insert create a new row in DB
func (u UserDAO) Insert(user *User) error {
	tx := u.db.Begin()

	if erro := tx.Create(user).Error; erro != nil {
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}

//Update alter a existent record
func (u UserDAO) Update(user *User) error {
	tx := u.db.Begin()

	if erro := tx.Save(user).Error; erro != nil {
		log.Println("Erro registro existente: ", erro)
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}

//Delete delete a row
func (u UserDAO) Delete(id int) error {
	tx := u.db.Begin()

	user := User{ID: id}
	if erro := tx.Delete(user).Error; erro != nil {
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}
