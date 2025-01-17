package config

type Config struct {
	Port int
	Host string
}

func New() Config {
	return Config{
		Port: 443,
		Host: "localhost",
	}
}
