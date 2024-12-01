package config

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type (
	Config struct {
		App      AppConfig
		Server   ServerConfig
		Database DatabaseConfig
		Security SecurityConfig
		Redis    RedisConfig
	}

	// AppConfig holds the configuration related to the application settings.
	AppConfig struct {
		Name        string
		Version     string
		Scheme      string
		Host        string
		Environment string
	}

	// ServerConfig holds the configuration for the server settings.
	ServerConfig struct {
		Port     string // The port on which the server will listen.
		Debug    bool   // Indicates if debug mode is enabled.
		TimeZone string // The time zone setting for the server.
	}

	// DatabaseConfig holds the configuration for the database connection.
	DatabaseConfig struct {
		Host     string // The host address of the database.
		Port     string // The port number of the database.
		DBName   string // The name of the database.
		UserName string // The username for connecting to the database.
		Password string // The password for connecting to the database.
		Debug    bool   // The Debug for debugging when query executed.
	}

	SecurityConfig struct {
		JWT JWTConfig
	}

	JWTConfig struct {
		Key             string
		ExpiredInSecond int64
		Label           string
	}

	RedisConfig struct {
		Addr     string
		Password string
	}
)

// LoadConfig loads the configuration from the specified filename.
func LoadConfig(filename string) (Config, error) {
	// Create a new Viper instance.
	v := viper.New()

	// Set the configuration file name, path, and environment variable settings.
	v.SetConfigName(fmt.Sprintf("config/%s", filename))
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file.
	if err := v.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file: %v\n", err)
		return Config{}, err
	}

	// Unmarshal the configuration into the Config struct.
	var config Config
	if err := v.Unmarshal(&config); err != nil {
		fmt.Printf("Error unmarshaling config: %v\n", err)
		return Config{}, err
	}

	return config, nil
}

// LoadConfigPath loads the configuration from the specified path.
func LoadConfigPath(path string) (Config, error) {
	// Create a new Viper instance.
	v := viper.New()

	// Set the configuration file name, path, and environment variable settings.
	v.SetConfigName(path)
	v.AddConfigPath(".")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Read the configuration file.
	if err := v.ReadInConfig(); err != nil {
		// Handle the case where the configuration file is not found.
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return Config{}, errors.New("config file not found")
		}
		return Config{}, err
	}

	// Parse the configuration into the Config struct.
	var c Config
	if err := v.Unmarshal(&c); err != nil {
		log.Printf("unable to decode into struct, %v", err)
		return Config{}, err
	}

	return c, nil
}
