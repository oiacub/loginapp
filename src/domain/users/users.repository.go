package users

type UserRepository interface {
	SaveUser(user User)
	UpdateUser(user User)
	GetUserByUsername(username string) *User
	GetUserById(id string) *User
}
