package config

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"io/ioutil"
)

var (
	ErrNoConfigFile = errors.New("Running without a config file")
	Cfg *Config
)

const (
	DefaultAppMode   string = "product"
	DefaultAddr      string = "127.0.0.1:8970"
	DefaultAccessLog string = "./logs/access.log"
	KB               int    = 1024
	MB               int    = KB * 1024
	GB               int    = MB * 1024
)

type Config struct {
	AppMode string `toml:"app_mode"`
	HttpAddr string `toml:"http_addr"`
	AccessLog string `toml:"access_log"`
	FileName string `toml:"-"`
	Redis RedisConfig `toml:"redis"`
	Database DBConfig `toml:"database"`
}

type RedisConfig struct {
	Host string `toml:"host"`
	Password string `toml:"password"`
}

type DBConfig struct {
	Type string `toml:"type"`
	User string `toml:"user"`
	Password string `toml:"password"`
	Host string `toml:"host"`
	DbName string `toml:"dbname"`
	TablePrefix string `toml:"tableprefix"`
}

func NewConfigDefault() *Config {
	cfg := new(Config)
	cfg.AppMode = DefaultAppMode
	cfg.HttpAddr = DefaultAddr
	// disable access log
	cfg.AccessLog = DefaultAccessLog
	return cfg
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
