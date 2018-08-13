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
}
