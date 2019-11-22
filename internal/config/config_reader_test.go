package config

import (
	"testing"
)

func TestReadConfig(t *testing.T) {
	configFile := "../../configs/config.yml"
	_, err := LoadConfig(configFile)

	if err != nil {
		t.Errorf("load config failed from %q, got err %s", configFile, err)
	}
}
