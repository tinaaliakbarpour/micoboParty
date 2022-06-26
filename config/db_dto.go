package config

type Database struct {
	Username string `yaml:"postgres.username"`
	Password string `yaml:"postgres.password"`
	Host     string `yaml:"postgres.host"`
	Port     string `yaml:"postgres.port"`
	DB       string `yaml:"postgres.db"`
	SslMode  string `yaml:"postgres.sslmode"`
	TimeZone string `yaml:"postgres.timezone"`
}
