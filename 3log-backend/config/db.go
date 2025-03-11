package config

type DB struct {
	Postgres Postgres
}

type Postgres struct {
	Host     string `mapstructure:"host" yaml:"host"`
	Port     string `mapstructure:"port" yaml:"port"`
	User     string `mapstructure:"user" yaml:"user"`
	Password string `mapstructure:"password" yaml:"password"`
	Dbname   string `mapstructure:"dbname" yaml:"dbname"`
	Sslmode  string `mapstructure:"sslmode" yaml:"sslmode"`
}
