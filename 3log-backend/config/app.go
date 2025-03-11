package config

type App struct {
	Name string `mapstructure:"name" yaml:"name"`
	Port string `mapstructure:"port" yaml:"port"`
}
