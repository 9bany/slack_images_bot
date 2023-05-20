package utils

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBSource string `mapstructure:"DB_SOURCE"`
	AppToken string `mapstructure:"SLACK_APP_TOKEN"`
	BotToken string `mapstructure:"SLACK_BOT_TOKEN"`
}

func LodConfig(path string, configName string) Config {
	var config Config
	viper.AddConfigPath(path)
	viper.SetConfigName(configName)
	viper.SetConfigType("env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
	viper.Unmarshal(&config)
	return config

}
