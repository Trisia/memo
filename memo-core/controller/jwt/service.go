package jwt

import "crypto/rand"

// JWT的HMAC密钥
var jwtKey = make([]byte, 16)

func init() {
	_, _ = rand.Reader.Read(jwtKey)
}

// New 创建新的Token
func New(c *Claims) string {
	return Create(c, jwtKey)
}

// Verify 验证Token，若token有效则返回Token中的信息
func Verify(token string) *Claims {
	return Validate(token, jwtKey)
}

// ResetKey 重置JWT密钥
func ResetKey() {
	_, _ = rand.Reader.Read(jwtKey)
}
