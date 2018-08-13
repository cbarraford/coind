package chaincfg

import (
	"os"

	"github.com/BurntSushi/toml"
)

type config struct {
	ChainParams ParamsSet
}

// Reads info from config file
func ReadConfig(configFile string) (ParamsSet, error) {
	cfg := config{
		ChainParams: DefaultParamSet,
	}

	_, err := os.Stat(configFile)
	if err != nil {
		return cfg.ChainParams, err
	}

	if _, err := toml.DecodeFile(configFile, &cfg); err != nil {
		return cfg.ChainParams, err
	}

	return cfg.ChainParams, nil
}
