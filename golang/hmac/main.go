package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"log"
)

func main() {
	hmacObject := hmac.New(sha256.New, []byte("fake_key"))
	hmacObject.Write([]byte("fake_message"))
	sig := hex.EncodeToString(hmacObject.Sum(nil))
	// a5bfb7421f682def4b577d43fe3ef9efab141d192c4fc13df6f325d79e58434d
	log.Println(sig)
}
