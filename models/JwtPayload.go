package models

import (
	"github.com/dgrijalva/jwt-go"
)

// JwtPayload - interface of the payloadjson
type JwtPayload struct {
	AccessToken  jwt.MapClaims
	RefreshToken jwt.MapClaims
}

type JwtPostBody struct {
	Name  string
	Email string
}
