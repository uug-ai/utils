package string

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const letterBytes = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charset = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz!@#$%^&*()-_=+[]{}|;:,.<>?/~`"

func GenerateShortLink() string {
	return RandStringBytesRmndr(6)
}

func RandStringBytesRmndr(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func RandKey() (string, error) {
	b := make([]byte, 32)
	_, err := cryptorand.Read(b)
	if err != nil {
		return "", fmt.Errorf("failed to generate secure random key: %v", err)
	}
	for i, v := range b {
		b[i] = charset[v%byte(len(charset))]
	}
	return string(b), nil
}

func GenerateKey(keyType string) (string, error) {
	switch keyType {
	case "public":
		key := "UUG" + RandStringBytesRmndr(16)
		return key, nil
	case "private":
		return RandKey()
	default:
		return "", fmt.Errorf("invalid key type: %s", keyType)
	}
}
