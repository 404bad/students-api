package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
		Address     string `yaml:"address"`
		Timeout     string `yaml:"timeout"`
		IdleTimeout string `yaml:"idle_timeout"`
} 

// env-default:"production" this are called struct tags	
type Config struct {
	Env string `yaml:"env" env:"ENV" env-required:"true"`
	HTTPServer HTTPServer `yaml:"http_server"`
	StoragePath string `yaml:"storage_path" env-required:"true"`
}

// in go we dont return erro from the function with NAMING MUSTLOADCONFIG, we will log the error and exit the application if there is an issue with loading the configuration. This is a common pattern in Go for functions that are expected to succeed and where failure is considered a fatal error.
func MustLoadConfig() *Config {
	var configPath string

	configPath = os.Getenv("CONFIG_PATH")

	// If the CONFIG_PATH environment variable is not set, we can also allow the user to specify the config path via command-line flags.
	if configPath == "" {
		configFlag := flag.String("config","", "path to the configuration file")
		flag.Parse()
		configPath = *configFlag
		// If the config path is still empty after checking both environment variable and command-line flag, we should log a fatal error and exit the application.
		if configPath == "" {
			log.Fatal("config path is not provided")
		}
	}
	// Check if the configuration file exists at the specified path
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist at path: %s", configPath)
	}

	// Load the configuration from the specified path
	var cfg	 Config
	
	// Use the cleanenv package to read the configuration file and populate the Config struct. If there is an error during this process, log a fatal error and exit the application.
	// If there is an error reading the configuration, log a fatal error and exit the application. This ensures that the application does not run with an invalid or missing configuration.
	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Failed to read configuration: %v", err)
	}

	return &cfg


}	