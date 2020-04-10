package utils

import (
	"net/http"
	"time"

	models "../models"
)

var (
	jwtPayload models.JwtPayload
	jwtExp     = 5 // in minutes
)

// GeneratePayload - shall generate token
func GeneratePayload(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	jwtPayload.AccessToken = map[string]string{
		"aud": "http://api.example.com",
		"iss": "https://krakend.io",
		"sub": "1234567890qwertyuio",
		"jti": "mnb23vcsrt756yuiomnbvcx98ertyuiop",
	}

	jwtPayload.RefreshToken = map[string]string{
		"aud": "http://api.example.com",
		"iss": "https://krakend.io",
		"sub": "1234567890qwertyuio",
		"jti": "mnb23vcsrt756yuiomn12876bvcx98ertyuiop",
	}

	jwtPayload.Exp = time.Now().Add(time.Duration(jwtExp)*time.Minute).UnixNano() / int64(time.Millisecond)

	responseJSON(w, jwtPayload)
}
