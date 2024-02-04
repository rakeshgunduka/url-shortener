// config/config.go
package config

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config/")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../../config/")
	viper.AutomaticEnv()

	profile := viper.GetString("PROFILE")
	if profile == "" {
		profile = "dev" // Default to dev environment if not set
		godotenv.Load("./../../config/dev.env")
		godotenv.Load("./config/dev.env")
	}

	viper.SetConfigName(profile)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	fmt.Printf("Using %s environment configuration\n", profile)
}

// GetConfigValue retrieves a configuration value by key.
func GetConfigValue(key string) string {
	return viper.GetString(key)
}

// GetStringMapString returns the value associated with the key as a map of strings.
func GetStringMapString(key string) map[string]string {
	return viper.GetStringMapString(key)
}

// GetPort retrieves the port configuration.
func GetPort() string {
	return viper.GetString("port")
}
