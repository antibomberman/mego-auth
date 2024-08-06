package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"log"
)

type Config struct {
	RedisHost string `env:"REDIS_HOST" required:"true"`
	RedisPort string `env:"REDIS_PORT" required:"true"`

	JWTSecret             string `env:"JWT_SECRET" required:"true"`
	AuthServiceServerPort string `env:"AUTH_SERVICE_SERVER_PORT" required:"true"`
	UserServiceAddress    string `env:"USER_SERVICE_ADDRESS" required:"true"`
	PostServiceAddress    string `env:"POST_SERVICE_ADDRESS" required:"true"`
	StorageServiceAddress string `env:"STORAGE_SERVICE_ADDRESS" required:"true"`
}

func Load() *Config {
	path := "./.env"
	var cfg Config
	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("failed to read config: %v", err)
	}
	return &cfg
}
