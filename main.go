package main

func main() {

	db, error := config.InitDB()
	if error != nil {
		panic(error)
	}

	config.InitHandle(db)
}
