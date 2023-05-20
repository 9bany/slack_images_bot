package utils

import (
	"os"
)

type Config struct {
	DBDriver string `mapstructure:"DB_DRIVER"`
	DBName   string `mapstructure:"DB_NAME"`
	DBSource string `mapstructure:"DB_SOURCE"`
	AppToken string `mapstructure:"SLACK_APP_TOKEN"`
	BotToken string `mapstructure:"SLACK_BOT_TOKEN"`
}

func LodConfig(path string) Config {
	return Config{
		DBDriver: os.Getenv("DB_DRIVER"),
		DBName:   os.Getenv("DB_NAME"),
		DBSource: os.Getenv("DB_SOURCE"),
		AppToken: os.Getenv("SLACK_APP_TOKEN"),
		BotToken: os.Getenv("SLACK_BOT_TOKEN"),
	}

}
