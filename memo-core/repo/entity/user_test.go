package entity

import (
	"fmt"
	"testing"
)

func TestUser_MarshalJSON(t *testing.T) {
	u := &User{Username: "AA", Password: "123", Salt: "123"}
	b, err := u.MarshalJSON()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("%s\n", b)
}
