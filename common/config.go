package common

import (
	"encoding/json"
	"os"
)

type Config struct {
	IP                    string `json:"IP"`
	Port                  int    `json:"port"`
	DBUser                string `json:"dbuser"`
	DBPassword            string `json:"dbpassword"`
	DBName                string `json:"dbname"`
	QiNiuScope            string `json:"qiniuscope"`
	QiNiuTokenExpiredTime int    `json:"qiniutokenexpiredtime"`
	QiNiuAkey             string `json:"qiniuakey"`
	QiNiuSkey             string `json:"qiniuskey"`
}

func NewConfig(conf_path string) *Config {
	// open config
	f, err := os.Open(conf_path)
	if err != nil {
		panic(err)
	}

	// decode config
	decoder := json.NewDecoder(f)
	conf := new(Config)
	err = decoder.Decode(conf)
	if err != nil {
		panic(err)
	}

	return conf
}
