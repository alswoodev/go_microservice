package utils

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
    hashed, err := bcrypt.GenerateFromPassword(
        []byte(password),
        bcrypt.DefaultCost,
    )
    return string(hashed), err
}

func Verify(hash string, password string) bool {
    return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}