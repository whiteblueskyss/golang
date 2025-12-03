package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type HTTPServer struct {
	Addr string `yaml:"address" env-required:"true"` // Addr is a string field with struct tags for YAML and environment variable requirements.
}

type Config struct {
	Env         string `yaml:"env" env:"ENV" env-required:"true"` // `yaml:"env"` indicates the YAML key, `env:"ENV"` indicates the environment variable name, and `env-required:"true"` means this field is mandatory. If not set, the program will log a fatal error.
	StoragePath string `yaml:"storage_path" env-required:"true"`
	HTTPServer  `yaml:"http_server"`
}

func MustLoad() *Config { // *config means it returns a pointer to Config struct.
	var configPath string

	configPath = os.Getenv("CONFIG_PATH") // Here os.Getenv fetches the value of the environment variable named "CONFIG_PATH". If this variable is not set, it returns an empty string.

	if configPath == "" {
		flags := flag.String("config", "", "path to the configuration file") // flag.String defines a string flag with specified name, default value, and usage string. It returns a pointer to a string variable that stores the value of the flag.
		flag.Parse()                                                         // parses means it reads the command-line arguments and assigns values to the defined flags.

		configPath = *flags

		if configPath == "" {
			log.Fatal("Config path is not set") // log.Fatal logs the message and then calls os.Exit(1), terminating the program.
		}
	}

	// fatal error means the program will log the error message and exit immediately.

	if _, err := os.Stat(configPath); os.IsNotExist(err) { // os.Stat returns file info. If the file does not exist, os.IsNotExist(err) returns true.
		log.Fatalf("config file does not exist: %s", configPath) // log.Fatalf formats according to a format specifier and logs the message, then exits.
	}

	var cfg Config

	err := cleanenv.ReadConfig(configPath, &cfg) // cleanenv is a package that reads configuration from files and environment variables into structs. Here, it reads the config file at configPath and populates the cfg struct.
	if err != nil {
		log.Fatalf("can not read config file: %s", err.Error())
	}

	return &cfg
}
