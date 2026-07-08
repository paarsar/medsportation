package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Port              string `mapstructure:"PORT"`
	DatabasePath      string `mapstructure:"DATABASE_PATH"`
	JwtSecret         string `mapstructure:"JWT_SECRET"`
	ResendApiKey      string `mapstructure:"RESEND_API_KEY"`
	NotificationEmail string `mapstructure:"NOTIFICATION_EMAIL"`
}

func LoadConfig() (config Config, err error) {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	// Handle standard environment variable mapping
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Default values
	viper.SetDefault("PORT", "8080")
	viper.SetDefault("DATABASE_PATH", "medsportation.db")
	viper.SetDefault("JWT_SECRET", "super-secret-change-me-in-prod")
	viper.SetDefault("NOTIFICATION_EMAIL", "info@medsportationlogistics.com")

	// Explicitly bind environment variables
	viper.BindEnv("PORT")
	viper.BindEnv("DATABASE_PATH")
	viper.BindEnv("JWT_SECRET")
	viper.BindEnv("RESEND_API_KEY")
	viper.BindEnv("NOTIFICATION_EMAIL")

	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			log.Println("No .env file found, using environment variables")
		} else {
			return
		}
	}

	err = viper.Unmarshal(&config)
	return
}
