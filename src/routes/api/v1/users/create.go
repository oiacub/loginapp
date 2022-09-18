package users

import (
	"encoding/json"
	"gotestweb/src/domain/tokens"
	"gotestweb/src/domain/users"
	"gotestweb/src/utils"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersUseCase := users.NewUserService()
	userFound, err := usersUseCase.GetUserByUsername(user.Username)

	if userFound.Id == "" {
		userSaved := usersUseCase.SaveUser(user)
		tokensUseCase := tokens.NewTokenService()
		token := tokensUseCase.GenerateNewToken(userSaved.Id)
		utils.SendResponse(w, token.Token, http.StatusOK)
	} else {
		utils.SendResponse(w, users.USER_EXISTS, http.StatusConflict)
	}

}
