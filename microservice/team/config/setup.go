package config

import (
	"github.com/spf13/viper"
	"path/filepath"
)

type Config struct {
	ExpireTime uint32 `mapstructure:"expire_time"`
	Issuer     string `mapstructure:"issuer"`
}

var Cfg = SetUp("../conf/config.yaml")

func SetUp(path string) Config {
	if path == "" {
		return Config{}
	}
	viper.SetConfigName(filepath.Base(path))
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)
	
	if err := viper.ReadInConfig(); err != nil {
		panic("fail to read configuration")
	}
	config := Config{}
	if err := viper.Unmarshal(&config); err != nil {
		panic("fail to unmarshal configuration")
	}
	return config
}
