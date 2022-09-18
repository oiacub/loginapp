package users

import (
	"errors"

	"github.com/google/uuid"
)

type UserService struct {
	usersRepository UserRepository
}

func (us *UserService) SaveUser(user User) User {
	uuid := uuid.New().String()
	user.Id = uuid
	us.usersRepository.SaveUser(user)
	return user
}

func (us *UserService) UpdateUser(oldUser User, updatedUser User) error {
	if oldUser.IsGoogleUser {
		updatedUser.Username = oldUser.Username
	}
	if oldUser.Username != updatedUser.Username {
		if us.CheckExistingUser(updatedUser.Username) {
			return errors.New(USER_EXISTS)
		}
	}
	us.usersRepository.UpdateUser(updatedUser)
	return nil
}

func (us *UserService) GetUserByUsername(username string) (User, error) {
	user := us.usersRepository.GetUserByUsername(username)
	return *user, nil
}

func (us *UserService) GetUserById(username string) (User, error) {
	user := us.usersRepository.GetUserById(username)
	return *user, nil
}

func (us *UserService) ValidateUser(username string, Password string) error {
	user := us.usersRepository.GetUserByUsername(username)
	if user.Id == "" {
		return errors.New(USERNAME_INVALID)
	}
	if user.Password != Password {
		return errors.New(USER_PASSWORD_INVALID)
	}
	return nil
}

func (us *UserService) CheckExistingUser(username string) bool {
	user := us.usersRepository.GetUserByUsername(username)
	if user.Username != "" {
		return true
	}
	return false
}

func (us *UserService) UseWithMySQLRepository() {
	us.usersRepository = &UserMySQLRepository{}
}

func (us *UserService) UseWithMemoryRepository() {
	us.usersRepository = &UserMemoryRepository{}
}

func NewUserService() *UserService {
	us := &UserService{}
	us.UseWithMySQLRepository()
	return us
}
