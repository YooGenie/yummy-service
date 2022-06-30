package common

import (
	"fmt"
	"testing"
)

func Test_Encrypt(t *testing.T) {

	key := "UjXn2r5u8x/A?D(G+KbPeSgVkYp3s6v9"
	plaintext1 := "01000000000"
	foo := Encrypt(key, plaintext1)
	fmt.Println("암호화 : ",foo)
	fmt.Println("암화화 풀기",Decrypt(key, "DxLjRSN17TPi1bc="))
}