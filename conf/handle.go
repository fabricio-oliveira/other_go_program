package conf

import (
	"fmt"
	"log"
	"net/http"

	co "github.com/fabricio-oliveira/simple-api/controller"

	"github.com/go-zoo/bone"
	"github.com/jinzhu/gorm"
)

//InitHandle init handlels
func InitHandle(db *gorm.DB) {

	porta := ":8081"
	portaStatic := ":8080"
	fmt.Println("WebServer go iniciado na porta ", porta)
	fmt.Println("WebServer static iniciado na porta ", portaStatic)

	mux := bone.New()

	user := co.NewUser(db)
	registerController(mux, user)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", http.StripPrefix("/", fs))

	// start golang server
	log.Fatal(http.ListenAndServe(porta, mux))

}

func registerController(mux *bone.Mux, c co.Rest) {
	mux.Get(c.URL(), http.HandlerFunc(c.Get))
	mux.Post(c.URL(), http.HandlerFunc(c.Post))
	mux.Put(c.URL(), http.HandlerFunc(c.Put))
	mux.Delete(c.URL(), http.HandlerFunc(c.Delete))
}
