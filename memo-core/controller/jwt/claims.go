package jwt

import (
	"bytes"
	"crypto/hmac"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/tjfoc/gmsm/sm3"
	"memo-core/setting"
	"strings"
	"time"
)

const (
	DefaultTokenDuration = 8 * time.Hour // Token 有效时长
	// Encode: {"alg":"HMAC-SM3","typ":"JWT"}
	header = "eyJhbGciOiJITUFDLVNNMyIsInR5cCI6IkpXVCJ9"
)

// Claims RFC 7519 JWT Claims
type Claims struct {
	Sub string `json:"sub"` // 用户名
	Typ int    `json:"typ"` // 用户类型 0 - 普通用户； 1 - 管理员; 2 - 应用
	Iat int64  `json:"iat"` // 生效时间，Unix 毫秒数
	Exp int64  `json:"exp"` // 过期时间，Unix 毫秒数
}

// Create 创建JWT
func Create(c *Claims, key []byte) string {
	start := time.Now()
	c.Iat = start.UnixMilli()
	c.Exp = start.Add(DefaultTokenDuration).UnixMilli()
	bodyBin, _ := json.Marshal(c)
	// base64url refer RFC 7514 Appendix C.
	payload := base64.RawURLEncoding.EncodeToString(bodyBin)

	h := hmac.New(sm3.New, key)
	_, _ = h.Write([]byte(header))
	_, _ = h.Write([]byte{'.'})
	_, _ = h.Write([]byte(payload))
	sig := base64.RawURLEncoding.EncodeToString(h.Sum(nil))
	return fmt.Sprintf("%s.%s.%s", header, payload, sig)
}

// Validate 验证JWT
func Validate(token string, key []byte) *Claims {
	start := strings.IndexByte(token, '.')
	end := strings.LastIndexByte(token, '.')

	if start == -1 || end == -1 || end == start {
		return nil
	}

	braw := token[start+1 : end]
	payload, err := base64.RawURLEncoding.DecodeString(braw)
	if err != nil {
		return nil
	}
	var c Claims
	err = json.Unmarshal(payload, &c)
	if err != nil {
		return nil
	}

	now := time.Now().UnixMilli()
	if now >= c.Exp {
		// 过期
		return nil
	}

	// 校验HMAC
	plaintext := []byte(token[:end])
	actual, err := base64.RawURLEncoding.DecodeString(token[end+1:])
	if err != nil {
		return nil
	}
	hash := hmac.New(sm3.New, key)
	hash.Write(plaintext)
	expect := hash.Sum(nil)

	// TODO: DEBUG 忽略校验
	if setting.Config != nil && setting.Config.Debug {
		return &c
	}

	if !bytes.Equal(expect, actual) {
		return nil
	}
	return &c
}
