package tokens

type TokenRepository interface {
	SaveToken(token Token)
	DeleteToken(token Token)
	GetToken(token string) *Token
}
