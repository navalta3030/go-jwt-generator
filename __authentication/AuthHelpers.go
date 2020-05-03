package authentication

import (
	"os"
)

var (
	// DummyClientID - test
	DummyClientID = "YOUR DUMMY ID"
	// DummyClientSecret - test
	DummyClientSecret = "YOUR DUMMY SECRET"
	// TestCallbackURL - test
	TestCallbackURL = "https://localhost:8020/login"
)

func getCallbackURL() string {

	url := os.Getenv("PRODUCTION_URL")

	if len(url) == 0 {
		return TestCallbackURL
	}

	return url

}

func getClientID() string {
	url := os.Getenv("CLIENT_ID")

	if len(url) == 0 {
		return DummyClientID
	}

	return url

}

func getClientSecret() string {
	url := os.Getenv("CLIENT_SECRET")

	if len(url) == 0 {
		return DummyClientSecret
	}

	return url

}
