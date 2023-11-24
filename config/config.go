package config

import (
	"log"

	env "github.com/Netflix/go-env"
)

var cfg Config

type Config struct {
	AppConfig      AppConfig
	DatabaseConfig DatabaseConfig
	RedisConfig    RedisConfig
	AuthConfig     AuthConfig
}

type AppConfig struct {
	AppHost      string `env:"APP_HOST"`
	Port         string `env:"APP_PORT,required=true"`
	WriteTimeout uint32 `env:"WRITE_TIMEOUT,required=true"` //8000 * time.Second,
	ReadTimeout  uint32 `env:"READ_TIMEOUT,required=true"`  //8000 * time.Second,
	IdleTimeout  uint32 `env:"IDLE_TIMEOUT,required=true"`  //8000 * time.Second,
}

type DatabaseConfig struct {
	DB_Host    string `env:"DB_HOST,required=true"`
	DB_User    string `env:"DB_USER,required=true"`
	DB_Pswd    string `env:"DB_PASSWORD,required=true"`
	DB_Name    string `env:"DB_NAME,required=true"`
	DB_Port    string `env:"DB_PORT,required=true"`
	DB_Sslmode string `env:"DB_SSLMODE,required=true"`
}

type RedisConfig struct {
	RedisPost     string `env:"REDIS_POST,required=true"`
	RedisPassword string `env:"REDIS_PASSWORD,required=true"`
	RedisDB       int    `env:"REDIS_DB,required=true"`
}
type AuthConfig struct {
	PrivateKey string `env:"PRIVATE_KEY,required=true"`
	PublicKey  string `env:"PUBLIC_KEY,required=true"`
}

func init() {
	_, err := env.UnmarshalFromEnviron(&cfg)
	if err != nil {
		log.Panic(err)
	}
}

func GetConfig() Config {
	return cfg
}
