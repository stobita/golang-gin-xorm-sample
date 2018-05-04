package models

type User struct {
	ID   int    `xorm:"'id'"`
	Name string `xorm:"'name'"`
}

func (u User) GetAll() *[]User {
	var users []User
	engine.Find(&users)
	return &users
}
