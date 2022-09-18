package tokens

import (
	"testing"
)

func TestToken_SaveToken(t *testing.T) {
	token := Token{
		UserId: "1234",
		Token:  "1234",
	}
	tokensService := TokenService{}
	tokensService.UseWithMemoryRepository()
	tokensService.SaveToken(token)
	token = tokensService.GetToken(token)
	if token.UserId == "" {
		t.Fatalf(`Error saving token`)
	}
}

func TestToken_DeleteToken(t *testing.T) {
	token := Token{
		UserId: "1234",
		Token:  "1234",
	}
	tokensService := TokenService{}
	tokensService.UseWithMemoryRepository()
	tokensService.SaveToken(token)
	tokensService.DeleteToken(token)
	token = tokensService.GetToken(token)
	if token.UserId != "" {
		t.Fatalf(`Error deleting token`)
	}
}

func TestToken_GenerateNewToken(t *testing.T) {
	tokensService := TokenService{}
	tokensService.UseWithMemoryRepository()
	token := tokensService.GenerateNewToken("1234")
	tokenRecover := tokensService.GetToken(token)
	if tokenRecover.UserId == "" {
		t.Fatalf(`Error generating new token`)
	}
}

func TestToken_TokenIsValid_ShouldBeValid(t *testing.T) {
	token := Token{
		UserId: "1234",
		Token:  "1234",
	}
	tokensService := TokenService{}
	tokensService.UseWithMemoryRepository()
	tokensService.SaveToken(token)
	if !tokensService.TokenIsValid("1234") {
		t.Fatalf(`Error token should be valid`)
	}
}

func TestToken_TokenIsValid_ShouldNotBeValid(t *testing.T) {
	token := Token{
		UserId: "1234",
		Token:  "1234",
	}
	tokensService := TokenService{}
	tokensService.UseWithMemoryRepository()
	tokensService.SaveToken(token)
	if tokensService.TokenIsValid("12345") {
		t.Fatalf(`Error token should be valid`)
	}
}
