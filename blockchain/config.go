package blockchain

import (
	"os"

	"github.com/BurntSushi/toml"
)

var cfg Configuration

type config struct {
	Blockchain Configuration
}

// Configuration - configuration settings for the wire package
type Configuration struct {
	BlockIdentifierHashFunction string
}

var DefaultConfiguration = Configuration{
	BlockIdentifierHashFunction: "BlockHash",
}

// Init - initialize the package configuration
func Init(configFile string) error {
	if configFile == "" {
		cfg = DefaultConfiguration
	}
	var err error
	cfg, err = ReadConfig(configFile)
	return err
}

// ReadConfig - reads the param set from a given configuration file location
func ReadConfig(configFile string) (Configuration, error) {
	cfg := config{
		Blockchain: DefaultConfiguration,
	}

	_, err := os.Stat(configFile)
	if err != nil {
		return cfg.Blockchain, err
	}

	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return cfg.Blockchain, err
	}

	return cfg.Blockchain, nil
}
