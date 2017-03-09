package models

//User date of User
type User struct {
	ID    int    `json:"id" db:"id"`
	Nome  string `json:"nome" db:"nome"`
	Idade int    `json:"idade" db:"idade"`
	Sexo  byte   `json:"sexo" db:"sexo"`
}
