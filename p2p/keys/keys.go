package keys

import (
	"crypto/sha256"
	"encoding/hex"
	"math/rand"
)

type privKey string
type pubKey string

func generatePhrase(n int) string {
	var letters = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	result := make([]byte, n)
	for x := range result {
		result[x] = letters[rand.Intn(len(letters))]
	}

	return string(result)
}

func GenerateKeys() (string, string) {
	randomPhrase := generatePhrase(10)
	privKey := sha256.Sum256([]byte(randomPhrase))
	pubKey := sha256.Sum256([]byte(hex.EncodeToString(privKey[:])))

	return hex.EncodeToString(privKey[:]), hex.EncodeToString(pubKey[:])
}

func Verify(privKey string, pubKey string) bool {
	hashToVerify := sha256.Sum256([]byte(privKey))
	return hex.EncodeToString(hashToVerify[:]) == pubKey
}