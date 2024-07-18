package lib

import (
	"crypto/sha1"
	"fmt"
)

func Hasing()  {
	var secret string = "secret"
	var sha = sha1.New()

	sha.Write([]byte(secret))
	encrypted := sha.Sum(nil)

	enc_string := fmt.Sprintf("%x", encrypted)

	fmt.Println(enc_string)
	fmt.Println(encrypted)
}