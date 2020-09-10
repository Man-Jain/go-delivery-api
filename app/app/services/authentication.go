package services

import (
	"app/app/models"
	"errors"
	"time"

	"app/app/controllers"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

// Claims struct is for JWT claims
type Claims struct {
	Email string
	jwt.StandardClaims
}

var jwtKey = []byte("asfa980asfdakjsdfksadf")

// LogIn check user credential and generate JWT token
func LogIn(email string, password string) (string, string, error) {
	var accessTokenString string
	var refreshToken string
	var err error

	user := models.User{}

	if result := controllers.DB.Where("email = ?", email).First(&user); result.Error != nil {
		return "", "", result.Error
	}

	if err = bcrypt.CompareHashAndPassword(user.Profile.Password, []byte(password)); err == nil {
		// Creating Access Token
		expirationTime := time.Now().Add(5 * time.Minute)

		claims := &Claims{
			Email: email,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTime.Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessTokenString, err = token.SignedString(jwtKey)

		// Creating Refresh Token

		expirationTimeRefresh := time.Now().Add(48 * time.Hour)

		claimsRefresh := &Claims{
			Email: email,
			StandardClaims: jwt.StandardClaims{
				// In JWT, the expiry time is expressed as unix milliseconds
				ExpiresAt: expirationTimeRefresh.Unix(),
			},
		}

		tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
		refreshToken, err = tokenR.SignedString(jwtKey)

		return accessTokenString, refreshToken, nil
	}

	return "", "", errors.New("Invalid Credentials")
}

// RefreshToken used for creating access token from refresh token
func RefreshToken(token string) (string, error) {
	var accessTokenString string
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}

		return "", err
	}
	if !tkn.Valid {
		return "", err
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()

	tokenS := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err = tokenS.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

func validateToken(token string) (string, error) {
	claims := Claims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return "", err
		}

		return "", err
	}
	if !tkn.Valid {
		return "", err
	}

	return claims.Email, nil
}
