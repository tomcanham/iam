package config

type Config struct {
	Port     int
	Host     string
	CertPath string
	KeyPath  string
}

func New() Config {
	return Config{
		Port:     443,
		Host:     "localhost",
		CertPath: "certs/server.crt",
		KeyPath:  "certs/server.key",
	}
}
