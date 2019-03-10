package main

import (
	"golang.org/x/crypto/bcrypt"
)

func main() {
	// p1 := "12345678"
	// _, _ := generatePassword(p1)
	// now, let's reverse it

	err := bcrypt.CompareHashAndPassword([]byte("$2a$08$RE5t7fZNLx39b0UmKomfN.Jc635sx3ipEbfbsdN/W8KRkx2KmP92S"), []byte("Adonese=1994"))
	if err != nil {
		panic(err)
	}

}

func generatePassword(pass string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(pass), 8)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}
