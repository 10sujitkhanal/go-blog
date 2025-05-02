package configs

import (
	"github.com/spf13/viper"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	ServerPort string
}

func LoadConfig() (*Config, error) {
	v := viper.New()
	v.SetConfigFile(".env") // You can use .env or .yaml
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	return &Config{
		DBHost:     v.GetString("DB_HOST"),
		DBPort:     v.GetInt("DB_PORT"),
		DBUser:     v.GetString("DB_USER"),
		DBPassword: v.GetString("DB_PASSWORD"),
		DBName:     v.GetString("DB_NAME"),
		ServerPort: v.GetString("SERVER_PORT"),
	}, nil
}
