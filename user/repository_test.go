package user

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInsertSucessful(t *testing.T) {
	db, error := InitDB(":memory:")
	if error != nil {
		t.Errorf("Problema abrir conexão banco %+v", error)
		return
	}
	defer db.Close()

	userRepository := newRepository(db)
	user := &Model{Name: "Fab", Age: 4, Sex: 1}
	if error = userRepository.insert(user); error != nil {
		t.Errorf("Problema na persistencia dos dados %+v", error)
		return
	}

	assert.Equal(t, error, nil)
	assert.Equal(t, 1, user.ID)
}

func TestInsertIdExitent(t *testing.T) {
	db, error := InitDB(":memory:")
	if error != nil {
		t.Errorf("Problema abrir conexão banco %+v", error)
		return
	}
	defer db.Close()

	userRepository := newRepository(db)
	user := &Model{ID: 1, Name: "Fab", Age: 4, Sex: 1}
	if error = userRepository.insert(user); error != nil {
		t.Errorf("Problema abrir conexão banco %+v", error)
		return
	}

	if error = userRepository.insert(user); error != nil {
		assert.Error(t, error, "UNIQUE constraint failed: users.id")
	}
}
