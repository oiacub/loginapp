package tokens

type TokenMemoryRepository struct {
	Tokens []Token
}

func (repo *TokenMemoryRepository) SaveToken(token Token) {
	repo.Tokens = append(repo.Tokens, token)
}

func (repo *TokenMemoryRepository) DeleteToken(token Token) {
	repo.Tokens = remove(repo.Tokens, repo.findIndexOfToken(token.Token))
}

func (repo TokenMemoryRepository) GetToken(token string) *Token {
	for _, v := range repo.Tokens {
		if token == v.Token {
			return &v
		}
	}
	return &Token{}
}

func (repo *TokenMemoryRepository) findIndexOfToken(token string) int {
	for i, v := range repo.Tokens {
		if v.Token == token {
			return i
		}
	}
	return -1
}

func remove(slice []Token, s int) []Token {
	return append(slice[:s], slice[s+1:]...)
}
