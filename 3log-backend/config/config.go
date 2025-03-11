package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App App `mapstructure:"app" yaml:"app"`
	Zap Zap `mapstructure:"zap" yaml:"zap"`
	DB  DB  `mapstructure:"db" yaml:"db"`
}

var C Config

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&C)
	if err != nil {
		panic(err)
	}
}
