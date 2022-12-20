package conf

import (
	"encoding/json"
	"os"
)

type Conf struct {
	ClientId string `json:"client"`
	Secret   string `json:"secret"`
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
