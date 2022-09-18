package tokens

import (
	"gotestweb/src/utils"
)

type TokenMySqlRepository struct {
}

func (repo *TokenMySqlRepository) SaveToken(token Token) {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	db.Exec("INSERT INTO tokens(token, userid) VALUES (?, ?)", token.Token, token.UserId)
	mysql.CloseConnection()
}

func (repo *TokenMySqlRepository) DeleteToken(token Token) {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	db.Exec("DELETE FROM tokens WHERE token=?", token.Token)
	mysql.CloseConnection()
}

func (repo TokenMySqlRepository) GetToken(token string) *Token {
	mysql := utils.MySqlDB{}
	db := mysql.GetConnection()
	tokenInfo := Token{
		Token: token,
	}
	db.QueryRow("SELECT userid FROM tokens WHERE token = ?", token).Scan(&tokenInfo.UserId)
	mysql.CloseConnection()
	return &tokenInfo
}
