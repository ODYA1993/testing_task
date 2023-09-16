package store

type ConfigDB struct {
	Host     string `env:"host" yaml:"host"`
	Port     string `env:"port" yaml:"port"`
	DBName   string `env:"db_name" yaml:"db_name"`
	UserName string `env:"user_name" yaml:"user_name"`
	Password string `env:"password" yaml:"password"`
	SslMode  string `env:"sslmode" yaml:"sslmode"`
}

func NewConfigDB() *ConfigDB {
	return &ConfigDB{}
}
