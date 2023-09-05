package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/febriyan1302/catalog-service/config"
	"github.com/febriyan1302/catalog-service/pkg/infrastructure/database"
	"github.com/febriyan1302/catalog-service/pkg/middleware"
	"github.com/febriyan1302/catalog-service/pkg/router"
	"github.com/gofiber/fiber/v2"
	"gopkg.in/yaml.v3"
)

type AppConfig struct {
	config *config.Config
}

// NewAppConfig initialize a new application configuration.
// Accept pointer of Config. It return pointer of AppConfig.
func NewAppConfig(config *config.Config) *AppConfig {
	return &AppConfig{
		config: config,
	}
}

// GetConfig get pointer of Config attribute.
func (cfg *AppConfig) GetConfig() *config.Config {
	return cfg.config
}

// loadConfig load configuration file based on environment. Accept environment string.
// It returns non-nil pointer of config and nil error if succeed, otherwise returns nil pointer of config and non-nil error.
func loadConfig(env string) (*config.Config, error) {
	var configFilePath string

	switch strings.ToLower(env) {
	case "production":
		configFilePath = "config/production.yaml"
	case "testing":
		configFilePath = "config/testing.yaml"
	default:
		configFilePath = "config/development.yaml"
	}

	configFile, err := os.Open(configFilePath)
	if err != nil {
		return nil, err
	}
	defer configFile.Close()

	var config config.Config
	decoder := yaml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		return nil, err
	}

	return &config, nil
}

func main() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development" // Default to development if APP_ENV is not set
	}

	config, err := loadConfig(env)
	if err != nil {
		panic(fmt.Sprintf("failed to load config: %s", err.Error()))
	}

	// Initializes database
	db, err := database.New().PGConnect(config)
	if err != nil {
		log.Fatal(err.Error())
	}

	app := fiber.New()

	jwt := middleware.NewAuthMiddleware("secret")

	router.RouteApi(app, db, &jwt)

	app.Listen(":" + string(config.App.RestPort))
}
