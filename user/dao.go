package user

import (
	"log"

	"github.com/jinzhu/gorm"
)

type userDAO struct {
	db *gorm.DB
}

func newUserDAO(db *gorm.DB) *userDAO {
	return &userDAO{db: db}
}

func (u userDAO) find(id int) (*User, error) {
	user := &User{}
	if erro := u.db.Where("id == ?", id).First(&user).Error; erro != nil {
		return nil, erro
	}
	return user, nil
}

func (u userDAO) insert(user *User) error {
	tx := u.db.Begin()

	if erro := tx.Create(user).Error; erro != nil {
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}

func (u userDAO) update(user *User) error {
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
func (u userDAO) delete(id int) error {
	tx := u.db.Begin()

	user := User{ID: id}
	if erro := tx.Delete(user).Error; erro != nil {
		tx.Rollback()
		return erro
	}

	tx.Commit()
	return nil
}
