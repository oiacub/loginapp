package tokens

import (
	"gotestweb/src/domain/users"

	"github.com/google/uuid"
)

type TokenService struct {
	tokenRepository TokenRepository
}

func (ts *TokenService) SaveToken(token Token) {
	ts.tokenRepository.SaveToken(token)
}

func (ts *TokenService) GetToken(token Token) Token {
	tokenInfo := ts.tokenRepository.GetToken(token.Token)
	return *tokenInfo
}

func (ts *TokenService) DeleteToken(token Token) error {
	ts.tokenRepository.DeleteToken(token)
	return nil
}

func (ts *TokenService) GenerateNewToken(userId string) Token {
	uuid := uuid.New().String()
	token := Token{UserId: userId, Token: uuid}
	ts.SaveToken(token)
	return token
}

func (ts *TokenService) GetUserByToken(token Token) (users.User, error) {
	tokenInfo := ts.tokenRepository.GetToken(token.Token)
	userRepo := users.UserMySQLRepository{}
	user := userRepo.GetUserById(tokenInfo.UserId)
	return *user, nil
}

func (ts *TokenService) TokenIsValid(token string) bool {
	if ts.tokenRepository.GetToken(token).UserId != "" {
		return true
	}
	return false
}

func (ts *TokenService) UseWithMySQLRepository() {
	ts.tokenRepository = &TokenMySqlRepository{}
}

func (ts *TokenService) UseWithMemoryRepository() {
	ts.tokenRepository = &TokenMemoryRepository{}
}

func NewTokenService() *TokenService {
	ts := &TokenService{}
	ts.UseWithMySQLRepository()
	return ts
}
