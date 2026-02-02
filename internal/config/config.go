package config

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Server   ServerConfig   `mapstructure:"server"`
	Database DatabaseConfig `mapstructure:"database"`
}

type ServerConfig struct {
	Port string `mapstructure:"port"`
	Mode string `mapstructure:"mode"`
}

type DatabaseConfig struct {
	DSN string `mapstructure:"dsn"`
}

var GlobalConfig *Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("./config") 

	// Environment variables
	// IRONOPS_SERVER_PORT=9090 ./main
	viper.SetEnvPrefix("IRONOPS")
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Warning: config file not found, using defaults or env vars: %v", err)
	}

	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode into struct: %v", err)
	}
    
    // Set defaults if needed
    if cfg.Server.Port == "" {
        cfg.Server.Port = "8080"
    }
    if cfg.Server.Mode == "" {
        cfg.Server.Mode = "debug"
    }

	GlobalConfig = &cfg
}
