package model

// Token token model
type Token struct {
	Key string
}

// Encrypt 生成token
func (t *Token) Encrypt(openID string, sessionKey string, expireTime int) (string, error) {
	return "", nil
}

// Decrypt 解析token
func (t *Token) Decrypt(token string) (openID string, sessionKey string, expireTime int) {
	return "", "", 0
}
