package models

import (
	"crypto/sha512"
	"database/sql"
	"encoding/base64"
)

type Session struct {
	ID        int64
	Key       string
	Value     string
	Token     string
	TokenHash string
}

type SessionRepo struct {
	DB *sql.DB
}

func (sr *SessionRepo) Create(token, key, value string) (*Session, error) {
	session := &Session{}
	session.Key = key
	session.Value = value
	session.Token = token
	tokenHash := sha512.Sum512([]byte(session.Token))
	session.TokenHash = base64.URLEncoding.EncodeToString(tokenHash[:])
	sr.deleteSession(token, key)
	row := sr.DB.QueryRow("INSERT INTO sessions (key, value, token_hash) VALUES ($1, $2, $3) RETURNING id",
		session.Key,
		session.Value,
		session.TokenHash)
	err := row.Scan(&session.ID)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (sr *SessionRepo) deleteSession(token, key string) error {
	tokenHash := sha512.Sum512([]byte(token))
	_, err := sr.DB.Exec(`DELETE FROM sessions WHERE token_hash = $1 AND key = $2`, tokenHash, key)
	if err != nil {
		return err
	}
	return nil
}

func (sr *SessionRepo) Find(token, key string) (*Session, error) {
	session := &Session{}
	tokenHash64 := sha512.Sum512([]byte(token))
	tokenHash := base64.URLEncoding.EncodeToString(tokenHash64[:])
	row := sr.DB.QueryRow(`SELECT id, token_hash, key, value FROM sessions WHERE token_hash = $1`, tokenHash)
	err := row.Err()
	if err != nil {
		return nil, err
	}
	row.Scan(&session.ID, &session.TokenHash, &session.Key, &session.Value)
	return session, nil
}
