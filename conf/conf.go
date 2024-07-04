package conf

import (
	"encoding/json"
	"fmt"
	"os"
)

type Conf struct {
	Spotify struct {
		ClientId    string `json:"client"`
		Secret      string `json:"secret"`
		RedirectUrl string `json:"redirectUrl"`
	} `json:"spotify"`
	DB struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
		Database string `json:"database"`
		SSLMode  string `json:"sslMode"`
	} `json:"db"`
}

func LoadConf(file *string) (*Conf, error) {
	f1, err := os.Open(*file)
	if err != nil {
		return nil, err
	}
	defer f1.Close()

	decoder := json.NewDecoder(f1)
	var conf Conf
	decoder.Decode(&conf)
	return &conf, nil
}

func (c *Conf) DBConfig() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		c.DB.Host, c.DB.Port, c.DB.User, c.DB.Password, c.DB.Database, c.DB.SSLMode)
}
