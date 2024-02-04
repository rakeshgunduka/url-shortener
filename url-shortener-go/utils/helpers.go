package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"os"
	"strings"
)

func Getenv(key string, defaultValue ...string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return key
	}
	return value
}

const base62Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func ConvertToBase62(input string) string {
	var result strings.Builder
	num := 0
	for _, char := range input {
		num = num*256 + int(char)
	}

	for num > 0 {
		remainder := num % 62
		num = num / 62
		result.WriteByte(base62Chars[remainder])
	}

	return result.String()
}

func GenerateHashWithSalt(data, salt string) string {
	hasher := sha256.New()
	hasher.Write([]byte(data + salt))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}
