package user

//Model date of User
type Model struct {
	ID   int    `json:"id" db:"id; primary_key:yes"`
	Name string `json:"name" db:"name"`
	Age  int    `json:"age" db:"age"`
	Sex  byte   `json:"sex" db:"sex"`
}

//TableName name model
func (Model) TableName() string {
	return "users"
}
