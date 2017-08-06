package main

import (
	"github.com/fabricio-oliveira/other_go_program/config"
)

func main() {

	db, error := config.InitDB()
	if error != nil {
		panic(error)
	}

	config.InitHandle(db)
}
