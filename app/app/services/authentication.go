package services

import (
	"app/app/models"
	"errors"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

var jwtKey = []byte("asfa980asfdakjsdfksadf")

// LogIn check user credential and generate JWT token
func LogIn(email string, password string) (string, string, error) {
	var accessTokenString string
	var refreshToken string
	var err error

	user := models.User{}

	if result := DB.Where("email = ?", email).First(&user); result.Error != nil {
		return "", "", result.Error
	}

	println(user.Profile.Email)

	if err = bcrypt.CompareHashAndPassword(user.Profile.Password, []byte(password)); err == nil {
		// Creating Access Token
		expirationTime := time.Now().Add(15 * time.Minute)

		claims := jwt.MapClaims{
			"ID":        user.ID,
			"Email":     user.Profile.Email,
			"ExpiresAt": expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessTokenString, err = token.SignedString(jwtKey)

		// Creating Refresh Token

		expirationTimeRefresh := time.Now().Add(48 * time.Hour)

		claimsRefresh := jwt.MapClaims{
			"ID":        user.ID,
			"Email":     user.Profile.Email,
			"ExpiresAt": expirationTimeRefresh.Unix(),
		}

		tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
		refreshToken, err = tokenR.SignedString(jwtKey)

		return accessTokenString, refreshToken, nil
	}

	return "", "", errors.New("Invalid Credentials")
}

// DeliveryLogIn check user credential and generate JWT token
func DeliveryLogIn(email string, password string) (string, string, error) {
	var accessTokenString string
	var refreshToken string
	var err error

	dp := models.DeliveryPerson{}

	if result := DB.Where("email = ?", email).First(&dp); result.Error != nil {
		return "", "", result.Error
	}

	println(dp.Profile.Email)

	if err = bcrypt.CompareHashAndPassword(dp.Profile.Password, []byte(password)); err == nil {
		// Creating Access Token
		expirationTime := time.Now().Add(15 * time.Minute)

		claims := jwt.MapClaims{
			"ID":        dp.ID,
			"Email":     dp.Profile.Email,
			"ExpiresAt": expirationTime.Unix(),
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		accessTokenString, err = token.SignedString(jwtKey)

		// Creating Refresh Token

		expirationTimeRefresh := time.Now().Add(48 * time.Hour)

		claimsRefresh := jwt.MapClaims{
			"ID":        dp.ID,
			"Email":     dp.Profile.Email,
			"ExpiresAt": expirationTimeRefresh.Unix(),
		}

		tokenR := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
		refreshToken, err = tokenR.SignedString(jwtKey)

		return accessTokenString, refreshToken, nil
	}

	return "", "", errors.New("Invalid Credentials")
}

// GetRefreshToken used for creating access token from refresh token
func GetRefreshToken(token string) (string, error) {
	var accessTokenString string
	claims := jwt.MapClaims{}

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
	claims["ExpiresAt"] = expirationTime.Unix()

	tokenS := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessTokenString, err = tokenS.SignedString(jwtKey)

	if err != nil {
		return "", err
	}

	return accessTokenString, nil
}

// ValidateToken validates token and returns email
func ValidateToken(token string) (uint, string, error) {
	claims := jwt.MapClaims{}

	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	println("error", err)

	if err != nil {
		return 9999, "", err
	}
	if !tkn.Valid {
		return 9999, "", err
	}

	return uint(claims["ID"].(float64)), claims["Email"].(string), nil
}

// ValidateUser validates token and returns email
func ValidateUser(token string, userType string) (bool, error) {
	id, userEmail, err := ValidateToken(token)
	if err != nil {
		return false, err
	}

	println("userEmail", userEmail)

	if userType == "root" {
		user := models.User{}

		if result := DB.Where("id = ? and root = 1", userEmail).First(&user); result.Error != nil {
			return false, result.Error
		}

		return true, nil
	} else if userType == "delivery_person" {
		dp := models.DeliveryPerson{}

		if result := DB.First(&dp, id); result.Error != nil {
			return false, result.Error
		}

		return true, nil
	} else {
		user := models.User{}

		if result := DB.First(&user, id); result.Error != nil {
			return false, result.Error
		}

		return true, nil
	}
}
