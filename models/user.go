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

func (u User) GetByID(id int) *User {
	user := User{ID: id}
	has, _ := engine.Get(&user)
	if has {
		return &user
	}
	return nil
}

func (u User) Insert(name string) *User {
	user := User{Name: name}
	_, err := engine.Insert(&user)
	if err == nil {
		return &user
	}
	return nil
}
