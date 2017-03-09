package main

import (
	"github.com/fabricio-oliveira/simple-api/conf"
)

func main() {

	db, error := conf.InitDB()
	if error != nil {
		panic(error)
	}

	conf.InitHandle(db)
}
