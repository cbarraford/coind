package chaincfg

import (
	"os"

	"github.com/BurntSushi/toml"
)

// Reads info from config file
func ReadConfig(configFile string) (ParamsSet, error) {
	set := DefaultParamSet

	_, err := os.Stat(configFile)
	if err != nil {
		return set, err
	}

	if _, err := toml.DecodeFile(configFile, &set); err != nil {
		return set, err
	}
	return set, nil
}
