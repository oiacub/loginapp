package users

type User struct {
	Id           string `json:"id"`
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Telephone    string `json:"telephone"`
	IsGoogleUser bool   `json:"isgoogleuser"`
}

type UserDTO struct {
	Fullname     string `json:"fullname"`
	Username     string `json:"username"`
	Telephone    string `json:"telephone"`
	IsGoogleUser bool   `json:"isgoogleuser"`
}

type GoogleUser struct {
	Id           string `json:"id"`
	Email        string `json:"email"`
	VerfiedEmail bool   `json:"verified_email"`
	Name         string `json:"name"`
	GivenName    string `json:"given_name"`
	FamilyName   string `json:"family_name"`
	Picture      string `json:"picture"`
	Hd           string `json:"hd"`
}

func (u User) ToDTO() UserDTO {
	return UserDTO{
		Fullname:     u.Fullname,
		Username:     u.Username,
		Telephone:    u.Telephone,
		IsGoogleUser: u.IsGoogleUser,
	}
}

func (u *User) UpdateFromDTO(dto UserDTO) {
	u.Username = dto.Username
	u.Fullname = dto.Fullname
	u.Telephone = dto.Telephone
}
