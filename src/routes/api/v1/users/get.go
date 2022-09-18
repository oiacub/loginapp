package users

import (
	"encoding/json"
	"gotestweb/src/domain/tokens"
	"gotestweb/src/utils"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	tokenService := tokens.NewTokenService()
	token := utils.GetToken(r)
	if tokenService.TokenIsValid(token) {
		user, _ := tokenService.GetUserByToken(tokens.Token{Token: token})
		jsonResp, _ := json.Marshal(user.ToDTO())
		utils.SendResponseJSON(w, jsonResp, http.StatusOK)
	} else {
		utils.SendResponse(w, "", http.StatusUnauthorized)
	}
}
