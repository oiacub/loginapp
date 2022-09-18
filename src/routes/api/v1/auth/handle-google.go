package auth

import (
	"gotestweb/src/domain/tokens"
	"gotestweb/src/domain/users"
	"gotestweb/src/utils"
	"net/http"

	"google.golang.org/api/oauth2/v2"
)

func HandleGoogleLoginExisting(w http.ResponseWriter, r *http.Request) {
	googleToken := utils.GetGoogleToken(r)
	tokenInfo, _ := verifyIdToken(googleToken)
	usersUseCase := users.NewUserService()
	user, _ := usersUseCase.GetUserByUsername(tokenInfo.Email)
	if user.Id != "" && user.IsGoogleUser {
		tokensUseCase := tokens.NewTokenService()
		token := tokensUseCase.GenerateNewToken(user.Id)
		utils.SendResponse(w, token.Token, http.StatusOK)
	} else {
		utils.SendResponse(w, users.USER_NOT_FOUND, http.StatusUnauthorized)
	}
}

func verifyIdToken(idToken string) (*oauth2.Tokeninfo, error) {
	oauth2Service, err := oauth2.New(&http.Client{})
	if err != nil {
		return nil, err
	}
	tokenInfoCall := oauth2Service.Tokeninfo()
	tokenInfoCall.IdToken(idToken)
	tokenInfo, err := tokenInfoCall.Do()
	if err != nil {
		return nil, err
	}
	return tokenInfo, nil
}

func HandleGoogleLogin(w http.ResponseWriter, r *http.Request) {
	googleToken := utils.GetGoogleToken(r)
	tokenInfo, _ := verifyIdToken(googleToken)
	usersUseCase := users.NewUserService()
	user, _ := usersUseCase.GetUserByUsername(tokenInfo.Email)
	if user.Id == "" {
		tokensUseCase := tokens.NewTokenService()
		newUser := usersUseCase.SaveUser(users.User{
			IsGoogleUser: true,
			Username:     tokenInfo.Email,
		})
		token := tokensUseCase.GenerateNewToken(newUser.Id)
		utils.SendResponse(w, token.Token, http.StatusOK)
	} else {
		utils.SendResponse(w, users.USER_EXISTS, http.StatusConflict)
	}
}
