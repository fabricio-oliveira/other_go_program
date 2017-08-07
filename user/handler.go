package user

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

//Handler is a controller for user
type Handler struct {
	user *repository
}

//NewHandler create a new user
func NewHandler(db *gorm.DB) *Handler {
	return &Handler{user: newRepository(db)}
}

func initHeader(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func valid(r *http.Request) bool {
	for _, key := range [3]string{"Name", "Age", "Sex"} {
		if val := bone.GetValue(r, key); val == "" {
			return false
		}
	}
	return true
}

//URL returnn URL to acess User
func (u Handler) URL() string {
	return "/user"
}

// Get return one User
func (u Handler) Get(w http.ResponseWriter, r *http.Request) {
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
	if err := json.NewEncoder(w).Encode(user); err != nil {
		log.Println("Erro Parse dados: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//Post Create one User
func (u Handler) Post(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	if valid(r) {
		log.Println("Erro parameter invalid")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &Model{}
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
	if err := json.NewEncoder(w).Encode(map[string]int{"id": user.ID}); err != nil {
		log.Println("Erro Parse dados: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

//Put update one User
func (u Handler) Put(w http.ResponseWriter, r *http.Request) {
	initHeader(w)

	id, erro := strconv.Atoi(bone.GetValue(r, "id"))
	if erro != nil {
		log.Println("Erro id inválido")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user := &Model{ID: id}
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
func (u Handler) Delete(w http.ResponseWriter, r *http.Request) {
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
