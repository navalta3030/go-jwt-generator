package routes

import (
	"encoding/json"
	"net/http"

	"go_jwt_generator/constants"
	"go_jwt_generator/controllers"
	"go_jwt_generator/models"
	"go_jwt_generator/utils"

	"github.com/dgrijalva/jwt-go"
)

var (
	jwtResponse  models.Jwt
	jwtPayload   models.JwtPayload
	jwtPostBody  models.JwtPostBody
	jwtExp       = 5 // in minutes
	jwtHeaderKid = "sim2"
	jwtSecret    = "secret" // encode this one into Base64 and change the static/auth.json k property

	logger   = utils.GeneralLogger
	logError = utils.ErrorLogger
)

// GenerateJWTFromPayload - shall generate token
func GenerateJWTFromPayload(w http.ResponseWriter, r *http.Request) {

	// get the body post message
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&jwtPostBody)
	if err != nil {
		logError.Println(err)
	}

	logger.Println("Accepting request from name: " + jwtPostBody.Name)
	jwtPayload.AccessToken = jwt.MapClaims{
		"name":  jwtPostBody.Name,
		"email": jwtPostBody.Email,
		"aud":   constants.Aud,
		"iss":   constants.Iss,
		"sub":   constants.Sub,
		"jti":   constants.Jti,
		"exp":   controllers.GetExpiration(),
	}

	jwtPayload.RefreshToken = jwt.MapClaims{
		"aud": constants.Aud,
		"iss": constants.Iss,
		"sub": constants.Sub,
		"jti": constants.Jti,
		"exp": controllers.GetExpiration(),
	}

	// Gather all objects for the JwtResponse
	jwtResponse.AccessToken = controllers.GetAccessToken(jwtPayload.AccessToken)
	jwtResponse.RefreshToken = controllers.GetRefreshToken(jwtPayload.RefreshToken)
	jwtResponse.Exp = controllers.GetExpiration()

	logger.Println("Returning request from: " + jwtPostBody.Name)

	//200
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Access-Control-Allow-Origin", "*")
	json.NewEncoder(w).Encode(&jwtResponse)
}
