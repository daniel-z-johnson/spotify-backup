package models

import (
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"github.com/daniel-z-johnson/spotify-backup/conf"
	"net/url"
)

const (
	oauthURL = "https://accounts.spotify.com/authorize"
)

type OAuth struct {
	DB   *sql.DB
	Conf *conf.Conf
}

func (oauth *OAuth) LinkForOAuth(challenge *string) (string, error) {
	authURL, err := url.Parse(oauthURL)
	if err != nil {
		return "", err
	}

	queries := authURL.Query()
	queries.Set("client_id", oauth.Conf.Spotify.ClientId)
	queries.Set("response_type", "code")
	queries.Set("redirect_uri", oauth.Conf.Spotify.RedirectUrl)
	queries.Set("scope", "playlist-read-private playlist-read-collaborative user-library-read")
	queries.Set("code_challenge_method", "S256")
	codehash := sha256.Sum256([]byte(*challenge))
	queries.Set("code_challenge", base64.URLEncoding.EncodeToString(codehash[:]))
	authURL.RawQuery = queries.Encode()

	return authURL.String(), nil
}
