// Copyright (c) 2016 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package chaincfg

import (
	"testing"

	"github.com/go-test/deep"
)

// TestMustRegisterPanic ensures that our sample btc config matches our default
// param set
func TestConfigMatchesDefault(t *testing.T) {
	t.Parallel()

	defaultParams := DefaultParamSet
	configParams, err := ReadConfig("sample-btc.conf")
	if err != nil {
		t.Errorf("Unable to parse config: %s", err)
	}

	if diff := deep.Equal(defaultParams, configParams); diff != nil {
		t.Error(diff)
	}

	configParams, err = ReadConfig("test/test1.conf")
	if err != nil {
		t.Errorf("Unable to parse config: %s", err)
	}
	// test that our config specification overrides the default
	if configParams.MainNetParams.Name != "test1" {
		t.Errorf("Name mismatch: got %s, expected 'test1'", configParams.MainNetParams.Name)
	}
	// test that if don't specify an attribute, we still get the default
	if configParams.MainNetParams.Net != 3652501241 {
		t.Errorf("Net mismatch: got %d, expected 3652501241", configParams.MainNetParams.Net)
	}
}
