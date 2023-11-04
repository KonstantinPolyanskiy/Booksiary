package password

import (
	"crypto/sha1"
	"fmt"
)

// Hash возвращает хеш пароля, хешируемый переданной солью
func Hash(password, salt string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
