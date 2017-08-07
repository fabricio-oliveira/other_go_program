package user

import (
	"log"

	"github.com/jinzhu/gorm"
)

type repository struct {
	db *gorm.DB
}

func newRepository(db *gorm.DB) *repository {
	return &repository{db: db}
}

func (u repository) find(id int) (*Model, error) {
	user := &Model{}
	if err := u.db.Where("id == ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return user, nil
}

func (u repository) insert(user *Model) error {
	tx := u.db.Begin()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (u repository) update(user *Model) error {
	tx := u.db.Begin()

	if err := tx.Save(user).Error; err != nil {
		log.Println("err existent record: ", err)
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

//Delete delete a row
func (u repository) delete(id int) error {
	tx := u.db.Begin()

	user := Model{ID: id}
	if err := tx.Delete(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}
