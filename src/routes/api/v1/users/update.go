package users

import (
	"encoding/json"
	"gotestweb/src/domain/tokens"
	"gotestweb/src/domain/users"
	"gotestweb/src/utils"
	"net/http"
)

func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	tokenService := tokens.NewTokenService()
	token := utils.GetToken(r)
	if tokenService.TokenIsValid(token) {
		user, _ := tokenService.GetUserByToken(tokens.Token{Token: token})
		userDTO := &users.UserDTO{}
		err := json.NewDecoder(r.Body).Decode(&userDTO)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		usersUseCase := users.NewUserService()
		errUpdate := usersUseCase.UpdateUser(user, users.User{
			Id:           user.Id,
			Fullname:     userDTO.Fullname,
			Username:     userDTO.Username,
			Password:     user.Password,
			Telephone:    userDTO.Telephone,
			IsGoogleUser: user.IsGoogleUser,
		})
		if errUpdate == nil {
			utils.SendResponse(w, "", http.StatusOK)
		} else {
			utils.SendResponse(w, errUpdate.Error(), http.StatusConflict)
		}
	} else {
		utils.SendResponse(w, "", http.StatusUnauthorized)
	}
}
