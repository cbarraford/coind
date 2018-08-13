package chainhash

import (
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	ChainHash Configuration
}

type Configuration struct {
	HashSize          int
	MaxHashStringSize int
}

// ReadConfig - reads the param set from a given configuration file location
func ReadConfig(configFile string) (Configuration, error) {
	cfg := config{
		ChainHash: DefaultConfiguration,
	}

	_, err := os.Stat(configFile)
	if err != nil {
		return cfg.ChainHash, err
	}

	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return cfg.ChainHash, err
	}

	return cfg.ChainHash, nil
}
