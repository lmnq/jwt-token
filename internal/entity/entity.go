package entity

import "time"

type AccessToken struct {
	Token string
}

type RefreshToken struct {
	GUID      string    `bson:"guid"`
	Hash      string    `bson:"hash"`
	ExpiresAt time.Time `bson:"expires_at"`
}

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}
