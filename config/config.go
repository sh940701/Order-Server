package config

import (
	"os"
	"github.com/naoina/toml"
)

type Config struct {
	Server struct {
		Mode string
		port string
	}

	DB map[string]map[string]interface{}
}

func GetConfig(fpath string) *Config {
	c := new(Config)
	if file, err := os.Open(fpath); err != nil {
		panic(err)
	} else {
		defer file.Close()
		if err := toml.NewDecoder(file).Decode(c); err != nil {
			panic(err)
		} else {
			return c
		}
	}
}