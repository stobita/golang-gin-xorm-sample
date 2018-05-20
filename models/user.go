package models

type User struct {
	ID       int    `xorm:"'id'"`
	Name     string `xorm:"'name'"`
	Password string `xorm:"'password'"`
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

func (u User) GetByName() *User {
	has, _ := engine.Get(&u)
	if has {
		return &u
	}
	return nil
}

func (u User) Insert() *User {
	_, err := engine.Insert(&u)
	if err == nil {
		return &u
	}
	return nil
}
