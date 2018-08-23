package wire

import (
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	Wire Configuration
}

// Configuration - configuration settings for the wire package
type Configuration struct {
	ProtocolVersion uint32
	MainNet         BitcoinNet
}

// ReadConfig - reads the param set from a given configuration file location
func ReadConfig(configFile string) (Configuration, error) {
	cfg := config{
		Wire: DefaultConfiguration,
	}

	_, err := os.Stat(configFile)
	if err != nil {
		return cfg.Wire, err
	}

	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return cfg.Wire, err
	}

	return cfg.Wire, nil
}
