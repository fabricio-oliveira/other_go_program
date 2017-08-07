package config

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fabricio-oliveira/other_go_program/user"
	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

//Rest Interface with methods mandatory
type Rest interface {
	URL() string
	Get(w http.ResponseWriter, r *http.Request)
	Post(w http.ResponseWriter, r *http.Request)
	Put(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

//InitHandle init handlels
func InitHandle(db *gorm.DB) {

	porta := ":8081"
	fmt.Println("WebServer go iniciado na porta ", porta)

	mux := bone.New()
	user := user.NewHandler(db)
	registerController(mux, user)

	// start golang server
	log.Fatal(http.ListenAndServe(porta, mux))
}

func registerController(mux *bone.Mux, c Rest) {
	mux.Get(c.URL()+":id", http.HandlerFunc(c.Get))
	mux.Post(c.URL(), http.HandlerFunc(c.Post))
	mux.Put(c.URL()+"/:id", http.HandlerFunc(c.Put))
	mux.Delete(c.URL()+"/:id", http.HandlerFunc(c.Delete))
}
