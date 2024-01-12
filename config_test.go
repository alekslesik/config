package config

import "testing"

func TestNewConfig(t *testing.T) {
	cfg, err := New()
	if err != nil {
		t.Error("error creating new config instance ")
	}

	if cfg == nil {
		t.Error("config is empty")
	}
}
