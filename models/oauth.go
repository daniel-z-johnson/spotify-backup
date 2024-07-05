package models

import (
	"database/sql"
	"fmt"
	"github.com/daniel-z-johnson/spotify-backup/conf"
	"net/url"
	"time"
)

const (
	oauthURL = "https://accounts.spotify.com/authorize"
	tokenURL = "https://accounts.spotify.com/api/token"
)

type OAuth struct {
	DB       *sql.DB
	Conf     *conf.Conf
	Sessions *SessionRepo
}

type TokenRequestBody struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	Code         string `json:"code"`
	RedirectURL  string `json:"redirect_url"`
	CodeVerifier string `json:"code_verifier"`
}

type Token struct {
	ID           int64
	TokenHash    string
	AccessToken  string
	TokenType    string
	Scope        string
	ExpiresIn    time.Time
	RefreshToken string
}

func (oauth *OAuth) LinkForOAuth(token *string) (string, error) {
	authURL, err := url.Parse(oauthURL)
	if err != nil {
		return "", err
	}

	queries := authURL.Query()
	queries.Set("client_id", oauth.Conf.Spotify.ClientId)
	queries.Set("response_type", "code")
	queries.Set("redirect_uri", oauth.Conf.Spotify.RedirectUrl)
	queries.Set("scope", "playlist-read-private playlist-read-collaborative user-library-read")
	queries.Set("show_dialog", "true")
	authURL.RawQuery = queries.Encode()

	return authURL.String(), nil
}

func (oauth *OAuth) GenTokenCode(code, challenge string) error {
	//tokenURL, err := url.Parse(tokenURL)
	//if err != nil {
	//	return err
	//}
	//requestBody := TokenRequestBody{
	//	ClientID:     oauth.Conf.Spotify.ClientId,
	//	Code:         code,
	//	RedirectURL:  oauth.Conf.Spotify.RedirectUrl,
	//	CodeVerifier: challenge,
	//	GrantType:    "authorization_code",
	//}

	return fmt.Errorf("Not implemented")
}
