package models

import "database/sql"

type Session struct {
	ID        int64
	key       string
	Token     string
	TokenHash string
}

type SessionRepo struct {
	db *sql.DB
}
