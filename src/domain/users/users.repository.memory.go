package users

type UserMemoryRepository struct {
	Users []User
}

func (repo *UserMemoryRepository) SaveUser(user User) {
	repo.Users = append(repo.Users, user)
}

func (repo *UserMemoryRepository) UpdateUser(user User) {
	repo.DeleteUser(user)
	repo.SaveUser(user)
}

func (repo UserMemoryRepository) GetUserByUsername(username string) *User {
	for _, v := range repo.Users {
		if username == v.Username {
			return &v
		}
	}
	return &User{}
}

func (repo UserMemoryRepository) GetUserById(id string) *User {
	for _, v := range repo.Users {
		if id == v.Id {
			return &v
		}
	}
	return &User{}
}

func (repo *UserMemoryRepository) findIndexOfUser(id string) int {
	for i, v := range repo.Users {
		if v.Id == id {
			return i
		}
	}
	return -1
}

func (repo *UserMemoryRepository) DeleteUser(user User) {
	repo.Users = remove(repo.Users, repo.findIndexOfUser(user.Id))
}

func remove(slice []User, s int) []User {
	return append(slice[:s], slice[s+1:]...)
}
