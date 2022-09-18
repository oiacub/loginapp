package users

import (
	"gotestweb/src/utils"
	"log"
)

type UserMySQLRepository struct {
}

func (repo *UserMySQLRepository) SaveUser(user User) {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	_, err := db.Exec("INSERT INTO users(id, username, telephone, fullname, password, isgoogleuser) VALUES (?, ?, ?, ?, ?, ?)", user.Id, user.Username, user.Telephone, user.Fullname, user.Password, user.IsGoogleUser)
	if err != nil {
		log.Panic(err)
	}
	mysql.CloseConnection()
}

func (repo *UserMySQLRepository) UpdateUser(user User) {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	db.Exec("UPDATE users SET username=?, telephone=?, fullname=?, password=?, isgoogleuser=? WHERE id=?", user.Username, user.Telephone, user.Fullname, user.Password, user.IsGoogleUser, user.Id)
	mysql.CloseConnection()
}

func (repo *UserMySQLRepository) GetUserByUsername(Username string) *User {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	var u User
	err := db.QueryRow("SELECT id, username, telephone, fullname, password, isgoogleuser FROM users WHERE username = ?", Username).Scan(&u.Id, &u.Username, &u.Telephone, &u.Fullname, &u.Password, &u.IsGoogleUser)
	mysql.CloseConnection()
	if err == nil {
		return &u
	} else {

		return &User{}
	}

}

func (repo *UserMySQLRepository) GetUserById(Id string) *User {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	var u User
	err := db.QueryRow("SELECT id, username, telephone, fullname, password, isgoogleuser FROM users WHERE Id = ?", Id).Scan(&u.Id, &u.Username, &u.Telephone, &u.Fullname, &u.Password, &u.IsGoogleUser)
	mysql.CloseConnection()

	if err == nil {
		return &u
	} else {
		return &User{}
	}
}
