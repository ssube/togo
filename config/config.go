package config

import (
	"io/ioutil"
	"log"
	"os/user"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Config struct {
	Default struct {
		Labels struct {
			Columns []string
			Sort    string
		}
		Projects struct {
			Columns []string
			Sort    string
		}
		Tasks struct {
			Columns []string
			Labels  []string
			Project string
			Sort    string
		}
	}
	Root  string
	Token string
}

func GetHome() string {
	usr, err := user.Current()

	if err != nil {
		log.Fatalf("error getting home dir: %s", err.Error())
	}

	return usr.HomeDir
}

func New(name string) *Config {
	path := filepath.Join(GetHome(), name)
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error loading config: %s", err.Error())
	}

	config := &Config{}
	err = yaml.Unmarshal(data, config)
	if err != nil {
		log.Fatalf("error parsing config: %s", err.Error())
	}

	return config
}
