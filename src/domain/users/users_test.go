package users

import (
	"testing"
)

func TestUser_NewUser(t *testing.T) {
	user := User{
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	usersService.SaveUser(user)
	_, err := usersService.GetUserByUsername(user.Username)
	if err != nil {
		t.Fatalf(`Error creating new user`)
	}
}

func TestUser_GetUserByUsername(t *testing.T) {
	user := User{
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	usersService.SaveUser(user)
	_, err := usersService.GetUserByUsername(user.Username)
	if err != nil {
		t.Fatalf(`Error getting user by username`)
	}
}

func TestUser_GetUserById(t *testing.T) {
	user := User{
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	userSaved := usersService.SaveUser(user)
	userSavedFromDb, _ := usersService.GetUserById(userSaved.Id)
	if userSavedFromDb.Id == "" {
		t.Fatalf(`Error getting user by id`)
	}
}

func TestUser_ValidateUserWrongPassword(t *testing.T) {
	user := User{
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
		Password:  "123",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	usersService.SaveUser(user)
	valUser := usersService.ValidateUser("oiacub@gmail.com", "1234")
	if valUser == nil {
		t.Fatalf(`Error validating user, should give an error`)
	}
}

func TestUser_ValidateUserCorrectPassword(t *testing.T) {
	user := User{
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
		Password:  "123",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	usersService.SaveUser(user)
	valUser := usersService.ValidateUser("oiacub@gmail.com", "123")
	if valUser != nil {
		t.Fatalf(`Error validating user, should NOT give an error`)
	}
}

func TestUser_UpdateUser(t *testing.T) {
	user := User{
		Id:        "A",
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	userSaved := usersService.SaveUser(user)
	userUpdated := User{
		Id:        userSaved.Id,
		Fullname:  "Octavio2",
		Telephone: "Iacub2",
		Username:  "oiacub@gmail.com2",
	}
	usersService.UpdateUser(user, userUpdated)
	userFromDb, _ := usersService.GetUserById(user.Id)
	if userFromDb.Fullname == userUpdated.Fullname {
		t.Fatalf(`Error updating user`)
	}
}

func TestUser_UpdateUserWithUsernameOfAnotherUser(t *testing.T) {
	user := User{
		Id:        "A",
		Fullname:  "Octavio",
		Telephone: "Iacub",
		Username:  "oiacub@gmail.com",
	}
	userTwo := User{
		Id:        "B",
		Fullname:  "Octavio2",
		Telephone: "Iacub2",
		Username:  "oiacub@gmail.com2",
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	userSaved := usersService.SaveUser(user)
	usersService.SaveUser(userTwo)
	userUpdated := User{
		Id:        userSaved.Id,
		Fullname:  "Octavio2",
		Telephone: "Iacub2",
		Username:  "oiacub@gmail.com2",
	}
	err := usersService.UpdateUser(user, userUpdated)
	if err == nil {
		t.Fatalf(`Error updating repeated user`)
	}
}

func TestUser_UpdateGoogleUserChangingEmail(t *testing.T) {
	user := User{
		Id:           "A",
		Fullname:     "Octavio",
		Telephone:    "Iacub",
		Username:     "oiacub@gmail.com",
		IsGoogleUser: true,
	}
	usersService := UserService{}
	usersService.UseWithMemoryRepository()
	userSaved := usersService.SaveUser(user)
	userUpdated := User{
		Id:           userSaved.Id,
		Fullname:     "Octavio2",
		Telephone:    "Iacub2",
		Username:     "oiacub@gmail.com2",
		IsGoogleUser: true,
	}
	usersService.UpdateUser(userSaved, userUpdated)
	userFromDb, _ := usersService.GetUserById(userSaved.Id)
	if userFromDb.Username != user.Username {
		t.Fatalf(`Error updating google user` + userFromDb.Username + " " + user.Username)
	}
}
