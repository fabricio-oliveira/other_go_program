package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

// User is a controller for user
type UserHandler struct {
	user *userRepository
}

func initHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

//NewUser create a new user
func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{user: newUserRepository(db)}
}

//URL returnn URL to acess User
func (u UserHandler) URL() string {
	return "/user/:id"
}

// Get return one User
func (u UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))

	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, erro := u.user.find(id)
	if erro != nil {
		log.Println("Erro: ", erro)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Body
	json.NewEncoder(w).Encode(user)
}

//Post Create one User
func (u UserHandler) Post(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &User{ID: id}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Erro Parse dados: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if erro := u.user.insert(user); erro != nil {
		log.Println("Erro registro existente: ", erro)
		w.WriteHeader(http.StatusConflict)
	}
	//Body
	w.WriteHeader(http.StatusAccepted)
}

//Put update one User
func (u UserHandler) Put(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &User{ID: id}
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		log.Println("Erro Parse dados: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if erro := u.user.update(user); erro != nil {
		w.WriteHeader(http.StatusConflict)
	}
	//Body
	w.WriteHeader(http.StatusAccepted)
}

//Delete remove one user
func (u UserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if erro := u.user.delete(id); erro != nil {
		log.Println("Erro ao excluir id: ", id, " ", erro)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//Status Code
	w.WriteHeader(http.StatusAccepted)
}
