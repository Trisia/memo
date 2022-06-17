package jwt

import (
	"testing"
)

func TestVerify(t *testing.T) {
	c := Claims{Sub: "user", Typ: 0}
	key := []byte{
		1, 2, 3, 4, 5, 6, 7, 8,
		1, 2, 3, 4, 5, 6, 7, 8,
	}

	jwt := Create(&c, key)

	cc := Validate(jwt, key)
	if cc == nil {
		t.Fatal("验证应该通过，但是不通过")
	}
}
