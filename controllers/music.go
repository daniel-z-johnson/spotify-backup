package controllers

import (
	"github.com/daniel-z-johnson/spotify-backup/models"
	"net/http"
)

type Music struct {
	Templates struct {
		OAuthPage Template
	}
	SessionsRepo *models.SessionRepo
}

func (m *Music) OauthPage(w http.ResponseWriter, r *http.Request) {

	m.Templates.OAuthPage.Execute(w, r, nil)
}
