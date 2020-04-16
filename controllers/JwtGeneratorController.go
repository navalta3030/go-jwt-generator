package controllers

import (
	"go_jwt_generator/constants"
	"go_jwt_generator/models"
	"go_jwt_generator/utils"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtResponse  models.Jwt
	jwtPayload   models.JwtPayload
	jwtExp       = 5 // in minutes
	jwtHeaderKid = constants.JwtHeaderKid
	jwtSecret    = constants.JwtSecret
	jwtPostBody  models.JwtPostBody

	logger   = utils.GeneralLogger
	logError = utils.ErrorLogger
)

// GetAccessToken -
func GetAccessToken(claims jwt.MapClaims) string {
	// Sign access_token with paload and Assign it to the final variable response
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken.Header["kid"] = jwtHeaderKid
	logger.Println(jwtSecret)
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		logError.Fatal(err)
	}
	return accessTokenString
}

//GetRefreshToken -
func GetRefreshToken(claims jwt.MapClaims) string {
	// Sign access_token with paload and Assign it to the final variable response
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken.Header["kid"] = jwtHeaderKid
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		logError.Fatal(err)
	}
	return refreshTokenString
}

// GetExpiration -
func GetExpiration() int64 {
	return time.Now().Add(time.Duration(jwtExp)*time.Minute).UnixNano() / int64(time.Millisecond)
}
