package models

// Jwt - response
type Jwt struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Exp          int64  `json:"expiration"`
}
