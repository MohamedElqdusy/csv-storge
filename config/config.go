package config

import (
	"csv-storage/utils"

	"github.com/kelseyhightower/envconfig"
)

type RedisConfig struct {
	RedisAdress   string `envconfig:"REDIS_ADRESS"`
	RedisPassword string `envconfig:"REDIS_PASSWORD"`
	RedisDataBase string `envconfig:"REDIS_DATABASE"`
	RedisPort     string `envconfig:"REDIS_PORT"`
}

func IniatilizeRedisConfig() *RedisConfig {
	var r RedisConfig
	err := envconfig.Process("", &r)
	utils.HandleError(err)
	return &r
}
