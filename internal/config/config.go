package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
	Db Database
}

type Database struct {
	Uri string `env:"DB_URI" env-required:"true"`
}

func MustLoadConfig() *Config {
	var cfg Config
	err := cleanenv.ReadEnv(&cfg)

	if err != nil {
		panic(err)
	}

	return &cfg
}
