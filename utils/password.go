package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(pass string) (string, error){
	hashed, err := bcrypt.GenerateFromPassword([]byte(pass), 14);

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

func CheckPasswordHash(password, hashedPass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password))

	return err == nil;
}