package models

import "errors"

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

const userSchema string = `CREATE TABLE users(
		id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
		username VARCHAR(30) NOT NULL,
		password VARCHAR(64) NOT NULL,
		email VARCHAR(40),
		created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP)`



type Users []User

var users = make(map[int]User)

func SetDefaultUser() {
	user := User{ID: 1, Username: "chris", Password: "admin123"}
	users[user.ID] = user
}

func GetUsers() Users {
	list := Users{}
	for _, user := range users {
		list = append(list, user)
	}
	return list
}

func GetUser(userId int) (User, error) {
	if user, ok := users[userId]; ok {
		return user, nil
	}
	return User{}, errors.New("El usuario no se encuentra dentro del mapa")
}

func SaveUser(user User) User {
	user.ID = len(users) + 1
	users[user.ID] = user
	return user
}

func UpdateUser(user User, username, password string) User {
	user.Username = username
	user.Password = password
	users[user.ID] = user
	return user
}

func DeleteUser(id int) {
	delete(users, id)
}
