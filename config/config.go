package config

import (
	"errors"
	"github.com/spf13/viper"
)

type Config struct {
	Server Server `yaml:"Server"`
	Logger Logger `yaml:"Logger"`
}

type Server struct {
	AppVersion  string `yaml:"AppVersion"`
	Port        string `yaml:"Port"`
	Environment string `yaml:"Environment"`
}

type Logger struct {
	Filename  string `yaml:"Filename"`
	MaxSize   string `yaml:"MaxSize"`
	LocalTime string `yaml:"LocalTime"`
	Compress  string `yaml:"Compress"`
}

var Configurations Config

func LoadConfig(fileName string) (*viper.Viper, error) {
	v := viper.New()

	v.SetConfigName(fileName)
	v.SetConfigType("yml")
	v.AddConfigPath(".")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {

		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, err
		}

	}

	return v, nil
}

func ParseConfig(v *viper.Viper) (*Config, error) {
	var c Config

	if err := v.Unmarshal(&c); err != nil {
		return nil, errors.New("unable to unmarshal config")
	}

	return &c, nil
}
