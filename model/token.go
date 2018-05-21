package model

// Token token model
type Token struct {
}

// Encrypt 生成token
func (t *Token) Encrypt(userID int, openID string, expireTime int) (string, error) {
	return "", nil
}

// Decrypt 解析token
func (t *Token) Decrypt(token string) (userID int, openID string, expireTime int) {
	return
}
