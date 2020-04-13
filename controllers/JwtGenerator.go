package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go_jwt_generator/constants"
	"go_jwt_generator/models"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtResponse  models.Jwt
	jwtPayload   models.JwtPayload
	jwtPostBody  models.JwtPostBody
	jwtExp       = 5 // in minutes
	jwtHeaderKid = "sim2"
	jwtSecret    = "secret" // encode this one into Base64 and change the static/auth.json k property
)

// GenerateJWTFromPayload - shall generate token
func GenerateJWTFromPayload(w http.ResponseWriter, r *http.Request) {

	// get the body post message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&jwtPostBody)
	if err != nil {
		panic(err)
	}

	//200
	w.WriteHeader(http.StatusOK)
	jwtPayload.AccessToken = jwt.MapClaims{
		"name":  jwtPostBody.Name,
		"email": jwtPostBody.Email,
		"aud":   constants.Aud,
		"iss":   constants.Iss,
		"sub":   constants.Sub,
		"jti":   constants.Jti,
		"exp":   getExpiration(),
	}

	jwtPayload.RefreshToken = jwt.MapClaims{
		"aud": constants.Aud,
		"iss": constants.Iss,
		"sub": constants.Sub,
		"jti": constants.Jti,
		"exp": getExpiration(),
	}

	// Gather all objects for the JwtResponse
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
