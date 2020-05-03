package constants

import (
	"os"
)

var (
	// Aud - Constant variable
	Aud = getEnvironmentVariable("JWT_AUDIENCE", "http://xray.marknavalta.com")
	// Iss - Constant variable
	Iss = getEnvironmentVariable("JWT_ISSUER", "https://xray.marknavalta.com")
	// Sub - Constant variable
	Sub = getEnvironmentVariable("JWT_SUBJECT", "information")
	// Jti - Constant variable
	Jti = getEnvironmentVariable("JWT_JTI", "mnb23vcsrt756yuiomnbvcx98ertyuiop")
	// Port - Constant variable
	Port = getEnvironmentVariable("JWT_PORT", ":8020")
	// JwtSecret - Constant variable
	JwtSecret = getEnvironmentVariable("JWT_SECRET", "secret") // encode this one into Base64 and change the static/auth.json k property
	// JwtHeaderKid - Constant variable
	JwtHeaderKid = getEnvironmentVariable("JWT_HEADER_KID", "sim2")
)

func getEnvironmentVariable(EnvName string, defaultValue string) string {
	value := os.Getenv(EnvName)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
