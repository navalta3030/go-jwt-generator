package utils

import (
	"log"
	"net/http"
	"time"

	models "go_jwt_generator/models"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtResponse  models.Jwt
	jwtPayload   models.JwtPayload
	jwtExp       = 5 // in minutes
	jwtHeaderKid = "sim2"
	jwtSecret    = "secret"
)

// GeneratePayload - shall generate token
func GeneratePayload(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	jwtPayload.AccessToken = jwt.MapClaims{
		"aud": "http://api.example.com",
		"iss": "https://krakend.io",
		"sub": "1234567890qwertyuio",
		"jti": "mnb23vcsrt756yuiomnbvcx98ertyuiop",
		"exp": getExpiration(),
	}

	jwtPayload.RefreshToken = jwt.MapClaims{
		"aud": "http://api.example.com",
		"iss": "https://krakend.io",
		"sub": "1234567890qwertyuio",
		"jti": "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
		"exp": getExpiration(),
	}

	// Get all the
	jwtResponse.AccessToken = getAccessToken(jwtPayload.AccessToken)
	jwtResponse.RefreshToken = getRefreshToken(jwtPayload.RefreshToken)
	jwtResponse.Exp = getExpiration()

	responseJSON(w, jwtResponse)
}

func getAccessToken(claims jwt.MapClaims) string {
	// Sign access_token with paload and Assign it to the final variable response
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken.Header["kid"] = jwtHeaderKid
	accessTokenString, err := accessToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal(err)
	}
	return accessTokenString
}

func getRefreshToken(claims jwt.MapClaims) string {
	// Sign access_token with paload and Assign it to the final variable response
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	refreshToken.Header["kid"] = jwtHeaderKid
	refreshTokenString, err := refreshToken.SignedString([]byte(jwtSecret))
	if err != nil {
		log.Fatal(err)
	}
	return refreshTokenString
}

func getExpiration() int64 {
	return time.Now().Add(time.Duration(jwtExp)*time.Minute).UnixNano() / int64(time.Millisecond)
}
