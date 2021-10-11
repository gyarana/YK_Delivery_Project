package model

type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type TokenIDs struct {
	UserID int
	UID    string
}
