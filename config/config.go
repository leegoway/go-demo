package config

import (
		"errors"
					"fmt"
	"io/ioutil"
	"github.com/BurntSushi/toml"
	)

var (
	ErrNoConfigFile = errors.New("Running without a config file")
)

const (
	DefaultAppMode string = "product"

	DefaultAddr string = "127.0.0.1:8970"

	DefaultAccessLog string = "./logs/access.log"

	KB int = 1024
	MB int = KB * 1024
	GB int = MB * 1024
)

type Config struct {

	// Addr can be empty to assign a local address dynamically
	AppMode string `toml:"app_mode"`

	HttpAddr string `toml:"http_addr"`

	AccessLog string `toml:"access_log"`

	FileName string `toml:"-"`
}


func NewConfigWithFile(fileName string) (*Config, error) {
	cfg := NewConfigDefault()

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	if err := toml.Unmarshal(data, cfg); err != nil {
		return nil, fmt.Errorf("newConfigwithData: unmarashal: %s", err)
	}
	cfg.FileName = fileName
	return cfg, err
}

func NewConfigDefault() *Config {
	cfg := new(Config)

	cfg.AppMode = DefaultAppMode
	cfg.HttpAddr = DefaultAddr

	// disable access log
	cfg.AccessLog = DefaultAccessLog

	return cfg
}
