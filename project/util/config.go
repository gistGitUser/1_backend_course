package util

import (
	"github.com/spf13/viper"
	"time"
)

// viper use mapstruct package for unmarshaling data
type Config struct {
	DBDriver            string        `mapstructure:"DB_DRIVER"`
	DBSource            string        `mapstructure:"DB_SOURCE"`
	ServerAddress       string        `mapstructure:"SERVER_ADDRESS"`
	TokenSymmetricKey   string        `mapstructure:"TOKEN_SYMMETRIC_KEY"`
	AccessTokenDuration time.Duration `mapstructure:"ACCESS_TOKEN_DURATION"`
}

func LoadConfig(path string) (config *Config, err error) {
	viper.AddConfigPath(path)  //file location
	viper.SetConfigName("app") //config name
	viper.SetConfigType("env") //type of config

	viper.AutomaticEnv() //override value if they exists

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	//unmarshal data to config
	err = viper.Unmarshal(&config)
	return
}
