package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type HttpConfig struct {
	Host string
	Port int
}

type SecurityConfig struct {
	ApiSign struct {
		AppKey    string `mapstructure:"app_key"`
		AppSecret string `mapstructure:"app_secret"`
	} `mapstructure:"api_sign"`
	Jwt struct {
		Key string `mapstructure:"key"`
	} `mapstructure:"jwt"`
}

type DbConfig struct {
	Type string
	Dsn  string
}
type RedisConfig struct {
	Addr         string
	Password     string
	DB           int
	ReadTimeout  float32
	WriteTimeout float32
}

type Log struct {
	Level      string
	Filename   string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
	Encoding   string
}

type Config struct {
	Env      string
	Http     HttpConfig     `mapstructure:"http"`
	Security SecurityConfig `mapstructure:"security"`
	Data     struct {
		Db    DbConfig    `mapstructure:"db"`
		Redis RedisConfig `mapstructure:"redis"`
	}
	Log Log `mapstructure:"log"`
}

func NewConfig(p string) *Config {
	envConf := os.Getenv("APP_CONF")
	if envConf == "" {
		envConf = p
	}
	fmt.Println("load conf file:", envConf)
	return getConfig(envConf)
}

func getConfig(path string) *Config {
	conf := viper.New()
	var config Config
	conf.SetConfigFile(path)
	err := conf.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = conf.Unmarshal(&config)
	if err != nil {
		panic(err)
	}
	return &config
}
