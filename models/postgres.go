package models

import (
	"database/sql"
	"fmt"
	"github.com/daniel-z-johnson/spotify-backup/conf"
	"github.com/daniel-z-johnson/spotify-backup/migrations"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/pressly/goose/v3"
)

func Open(conf *conf.Conf) (*sql.DB, error) {
	db, err := sql.Open("pgx", conf.DBConfig())
	if err != nil {
		return nil, err
	}
	err = migration(db)
	if err != nil {
		return nil, fmt.Errorf("issue with migrations: '%s'", err.Error())
	}
	return db, nil
}

func migration(db *sql.DB) error {
	err := goose.SetDialect("postgres")
	if err != nil {
		return fmt.Errorf("set Dialect Issue: '%s'", err.Error())
	}
	goose.SetBaseFS(migrations.Migrations)
	defer func() { goose.SetBaseFS(nil) }()
	err = goose.Up(db, ".")
	if err != nil {
		return fmt.Errorf("migration failed: '%s'", err.Error())
	}
	return nil
}
