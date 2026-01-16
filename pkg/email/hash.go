package email

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func Hash(str string) (string, error) {
	salt := make([]byte, 16)
	_, err := rand.Read(salt)
	if err != nil {
		fmt.Println("Ошибка при генерации соли")
		return "", err
	}

	saltHex := hex.EncodeToString(salt)
	
	strWithSalt := str + saltHex

	hash := sha256.Sum256([]byte(strWithSalt))
	hashString := hex.EncodeToString(hash[:])

	return hashString, nil
}