package models

type JwtPayload struct {
	AccessToken  map[string]string `json:"access_token"`
	RefreshToken map[string]string `json:"refresh_token"`
	Exp          int64
}
