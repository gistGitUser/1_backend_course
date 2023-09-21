package util

import "github.com/spf13/viper"

// viper use mapstruct package for unmarshaling data
type Config struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
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
