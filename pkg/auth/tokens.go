package auth

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

type SignedDetails struct {
	ID    string
	Email string
	jwt.StandardClaims
}

var secret = []byte("my_secret")

func TokenGenerator(email string, id string) (signedToken string, signedFreshToken string, err error) {
	claims := &SignedDetails{
		ID:    id,
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}

	refreshclaims := &SignedDetails{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(168)).Unix(),
		},
	}

	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(secret))

	if err != nil {
		log.Println("err with jwt generation")
		return
	}

	refreshtoken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshclaims).SignedString([]byte(secret))
	if err != nil {
		log.Println("err with jwt generation")
		return
	}

	return token, refreshtoken, err

}

func ValidateToken(signedtoken string) (claims *SignedDetails, err error) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		log.Println("the token is invalid")
		return nil, err
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, err
	}
	return claims, nil
}