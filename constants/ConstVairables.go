package constants

import (
	"os"
)

var (
	// Aud - Constant variable
	Aud = getEnvironmentVariable("JWT_AUDIENCE", "http://localhost")
	// Iss - Constant variable
	Iss = getEnvironmentVariable("JWT_ISSUER", "https://marknavalta.com")
	// Sub - Constant variable
	Sub = getEnvironmentVariable("JWT_SUBJECT", "1234567890qwertyuio")
	// Jti - Constant variable
	Jti = getEnvironmentVariable("JWT_JTI", "mnb23vcsrt756yuiomnbvcx98ertyuiop")
	// Port - Constant variable
	Port = getEnvironmentVariable("JWT_PORT", ":8020")
)

func getEnvironmentVariable(EnvName string, defaultValue string) string {
	value := os.Getenv(EnvName)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
