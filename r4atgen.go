package main

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
)

func main() {
	s := "test string"
	h := sha256.New()
	h.Write([]byte(s))

	sha256Hash := base64.URLEncoding.EncodeToString(h.Sum(nil))

	fmt.Println(s, sha256Hash)
}
