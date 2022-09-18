package auth

import (
	"encoding/json"
	"gotestweb/src/domain/tokens"
	"gotestweb/src/domain/users"
	"gotestweb/src/utils"
	"net/http"
)

func LoginUser(w http.ResponseWriter, r *http.Request) {
	user := users.User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	usersUseCase := users.NewUserService()
	validUser := usersUseCase.ValidateUser(user.Username, user.Password)
	if validUser == nil {
		user, err = usersUseCase.GetUserByUsername(user.Username)
		tokensUseCase := tokens.NewTokenService()
		token := tokensUseCase.GenerateNewToken(user.Id)
		utils.SendResponse(w, token.Token, http.StatusOK)
	} else {
		utils.SendResponse(w, validUser.Error(), http.StatusUnauthorized)
	}
}
