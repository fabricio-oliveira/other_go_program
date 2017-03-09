package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/fabricio-oliveira/simple-api/dao"
	"github.com/fabricio-oliveira/simple-api/models"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

// User is a controller for user
type User struct {
	userDAO *dao.User
}

func initHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//NewUser create a new user
func NewUser(db *gorm.DB) *User {
	return &User{userDAO: dao.NewUser(db)}
}


//URL returnn URL to acess User
func (u User) URL() string {
	return "/user/:id"
}

// Get return one User
func (u User) Get(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))

	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, erro := u.userDAO.Find(id)
	if erro != nil {
		log.Println("Erro: ", erro)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Body
	json.NewEncoder(w).Encode(user)
}

//Post Create one User
func (u User) Post(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &models.User{ID: id}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Erro Parse dados: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if erro := u.userDAO.Insert(user); erro != nil {
		log.Println("Erro registro existente: ", erro)
		w.WriteHeader(http.StatusConflict)
	}
	//Body
	w.WriteHeader(http.StatusAccepted)
}

//Put update one User
func (u User) Put(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &models.User{ID: id}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Erro Parse dados: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if erro := u.userDAO.Update(user); erro != nil {
		w.WriteHeader(http.StatusConflict)
	}
	//Body
	w.WriteHeader(http.StatusAccepted)
}

//Delete remove one user
func (u User) Delete(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if erro := u.userDAO.Delete(id); erro != nil {
		log.Println("Erro ao excluir id: ", id, " ", erro)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Status Code
	w.WriteHeader(http.StatusAccepted)
}
