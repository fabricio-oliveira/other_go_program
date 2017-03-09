package dao

import (
	"log"

	"github.com/fabricio-oliveira/simple-api/models"

	"github.com/jinzhu/gorm"
)

//User UserDAO
type User struct {
	db *gorm.DB
}

//NewUser create new use
func NewUser(db *gorm.DB) *User {
	return &User{db: db}
}

//Find return a user
func (u User) Find(id int) (*models.User, error) {
	user := &models.User{}
	if erro := u.db.Where("id == ?", id).First(&user).Error; erro != nil {
		return nil, erro
	}
	return user, nil
}

//Insert create a new row in DB
func (u User) Insert(user *models.User) error {
	tx := u.db.Begin()

	if erro := tx.Create(user).Error; erro != nil {
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}

//Update alter a existent record
func (u User) Update(user *models.User) error {
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
func (u User) Delete(id int) error {
	tx := u.db.Begin()

	user := models.User{ID: id}
	if erro := tx.Delete(user).Error; erro != nil {
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}
