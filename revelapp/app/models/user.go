package models

import "math/rand"

type User struct {
	Uid         int
	AccessToken string
	Sub					string
	Name				string
	Given_Name		string
	Family_Name	string
	Profile			string
	Picture			string
	Gender			string
}

var db = make(map[int]*User)

func GetUser(id int) *User {
	return db[id]
}

func NewUser() *User {
	user := &User{Uid: rand.Intn(10000)}
	db[user.Uid] = user
	return user
}
