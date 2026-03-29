package models

// Renew struct to describe refresh token object.
type Renew struct {
	RefreshToken string `json:"refresh_token"`
}

type Tokens struct {
	access  string `json:"access_token"`
	refresh string `json:"refresh_token"`
}
