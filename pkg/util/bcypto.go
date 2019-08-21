package util

import "golang.org/x/crypto/bcrypt"

func PasswordBCrypto(password string) string {
	cryptPwd, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(cryptPwd)
}

func VerifyPassword(hashPwd string, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(password))
	if err != nil{
		return false
	}
	return true
}
