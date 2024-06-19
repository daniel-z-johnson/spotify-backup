package models

import "database/sql"

type Session struct {
	ID        int64
	Key       string
	Token     string
	TokenHash string
}

type SessionRepo struct {
	db *sql.DB
}

func (sr *SessionRepo) Save(s *Session) (*Session, error) {
	row := sr.db.QueryRow(`INSERT INTO sessions (token_has, key) 
									VALUES ($1, $2) RETURNING id`,
		s.TokenHash, s.Key)
	err := row.Scan(&s.ID)
	if err != nil {
		return nil, err
	}
	return s, nil
}
