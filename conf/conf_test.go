package conf

import (
	"testing"
)

func TestLoadConf(t *testing.T) {
	fileLoc := "conf_test.json"
	c, e := LoadConf(&fileLoc)
	if e != nil {
		t.Fatal("Err should be nil")
	}
	if c == nil {
		t.Fatal("config loaded should not be nil")
	}
	if c.Spotify.ClientId != "tClient" {
		t.Error("client should have value tClient")
	}
	if c.Spotify.Secret != "tSecret" {
		t.Error("secret should have value tSecret")
	}
	if c.Spotify.RedirectUrl != "tURL" {
		t.Error("redirectUrl should have value tURL")
	}
}

func TestLoadConfDB(t *testing.T) {
	fileLoc := "conf_test.json"
	c, e := LoadConf(&fileLoc)
	if e != nil {
		t.Fatal("Err should be nil")
	}
	if c.DBConfig() != "host=host port=port user=user password=password dbname=database sslmode=sslMode" {
		t.Error("incorrect db conf string from config")
	}
}
