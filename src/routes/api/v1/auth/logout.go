package auth

import (
	"gotestweb/src/domain/tokens"
	"gotestweb/src/utils"
	"net/http"
)

func HandleLogout(w http.ResponseWriter, r *http.Request) {
	tokensUseCase := tokens.NewTokenService()
	token := utils.GetToken(r)
	tokensUseCase.DeleteToken(tokens.Token{
		Token: token,
	})
}
