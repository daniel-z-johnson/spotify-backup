package middleware

import (
	"fmt"
	"github.com/daniel-z-johnson/spotify-backup/context"
	"github.com/daniel-z-johnson/spotify-backup/models"
	"github.com/daniel-z-johnson/spotify-backup/random"
	"net/http"
)

const (
	cookieSession = "cookieSession"
)

type SessionStore struct {
	Session *models.SessionRepo
}

func (sstore *SessionStore) HandleState() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			token := context.GetSession(r.Context())

			code, err := sstore.Session.Find(*token, "code")
			if err != nil || code.ID == 0 {
				c, _ := random.SessionToken()
				code, _ = sstore.Session.Create(*token, "code", c)
			}
			cText := r.Context()
			cText = context.WithState(cText, &code.Value)
			r = r.WithContext(cText)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func Session() func(handler http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			s, err := readCookie(r, cookieSession)
			// cookie may not exist
			if err != nil {
				s, _ = random.SessionToken()
				setCookie(w, cookieSession, s)
			}
			c := r.Context()
			c = context.WithSession(c, &s)
			r = r.WithContext(c)
			next.ServeHTTP(w, r)
		}
		return http.HandlerFunc(fn)
	}
}

func setCookie(w http.ResponseWriter, name, value string) {
	cookie := newCookie(name, value)
	http.SetCookie(w, cookie)
}

func readCookie(r *http.Request, name string) (string, error) {
	c, err := r.Cookie(name)
	if err != nil {
		return "", fmt.Errorf("%s: %w", name, err)
	}
	return c.Value, nil
}

func deleteCookie(w http.ResponseWriter, name string) {
	cookie := newCookie(name, "")
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
}

func newCookie(name, value string) *http.Cookie {
	cookie := http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		HttpOnly: true,
	}
	return &cookie
}
