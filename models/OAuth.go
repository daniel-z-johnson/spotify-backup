package models

import "net/url"

const (
	AuthEndpoint string = "accounts.spotify.com"
	OAuthScheme  string = "https"
	// ContentType - OAuth docs for Spotify says to use form encoding
	ContentType string = "application/x-www-form-urlencoded"
	// Scope - needed to read user library
	Scope string = "user-library-read playlist-read-private playlist-read-collaborative"
)

type SpotifyOAuthResponse struct {
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	Scope        string `json:"scope"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type SpotifyOAuthService struct {
	ClientID     string
	ClientSecret string
	RedirectURL  string
	AccessToken  string
	RefreshToken string
}

func (spot *SpotifyOAuthService) SetTokenAndRefresh(accessToken, refreshToken string) {
	spot.AccessToken = accessToken
	spot.RefreshToken = refreshToken
}

func (spot *SpotifyOAuthService) GetOAuthURL() (string, error) {
	url := url.URL{}
	url.Scheme = OAuthScheme
	url.Host = AuthEndpoint
	url.Path = "authorize"
	v := url.Query()
	v.Set("scope", Scope)
	v.Set("client_id", spot.ClientID)
	v.Set("redirect_uri", spot.RedirectURL)
	url.RawQuery = v.Encode()
	return url.String(), nil
}
