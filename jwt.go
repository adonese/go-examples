package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenClaims struct {
	ServiceName string `json:"service_name"`
	jwt.StandardClaims
}


func main() {
key := []byte("my-token")
service_id := "morsal"
token, err := generateJWT(key, service_id)
if err != nil{
	panic(err)
}
claim, err := verifyJWT(token)
if err != nil{
	panic(err)
}
fmt.Println(claim)
}

func generateJWT(secretKey []byte, serviceID string) (string, error) {
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	expiresAt := time.Now().Add(time.Hour * 24).UTC().Unix()

	claims := TokenClaims{
		serviceID,
		jwt.StandardClaims{
			ExpiresAt: expiresAt,
			Issuer:    "noebs",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign and get the complete encoded token as a string using the secret
	if tokenString, err := token.SignedString(secretKey); err == nil {
		fmt.Println(tokenString)
		return tokenString, nil
	} else{
	return "", err
}
}

func verifyJWT(tokenString string) (jwt.Claims, error) {

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		} else {
			return []byte("my-token"), nil
		}
	})
	if err != nil {
		return nil, err
	}
	return token.Claims, err
}
