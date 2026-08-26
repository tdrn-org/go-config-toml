//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-config-toml"
)

func TestDefaults(t *testing.T) {
	cfg := &Config{}
	err := config.Defaults(defaultsData, cfg)
	require.NoError(t, err)
}

func TestLoadValid(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_valid.toml", defaultsData, cfg, true)
	require.NoError(t, err)
}

func TestLoadInvalidDuration(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_invalid_duration.toml", defaultsData, cfg, true)
	require.Error(t, err)
}

func TestLoadInvalidNetworks(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_invalid_networks.toml", defaultsData, cfg, true)
	require.Error(t, err)
}

func TestLoadInvalidRegexp(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_invalid_regexp.toml", defaultsData, cfg, true)
	require.Error(t, err)
}

func TestLoadInvalidTimeLocation(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_invalid_time_location.toml", defaultsData, cfg, true)
	require.Error(t, err)
}

func TestLoadInvalidURLs(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_invalid_urls.toml", defaultsData, cfg, true)
	require.Error(t, err)
}

func TestLoadInvalidUnknownField(t *testing.T) {
	cfg := &Config{}
	err := config.Load("testdata/config_invalid_unknown_field.toml", defaultsData, cfg, true)
	require.Error(t, err)
}

var defaultsData []byte = []byte(`
duration = ""
networks = [ ]
regexp = ""
time_location = ""
urls = [ "" ]
`)

type Config struct {
	Duration     config.DurationSpec     `toml:"duration,omitempty"`
	Networks     config.NetworkSpecs     `toml:"networks,omitempty"`
	TimeLocation config.TimeLocationSpec `toml:"time_location,omitempty"`
	URLs         config.URLSpecs         `toml:"urls,omitempty"`
	Regexp       config.RegexpSpec       `toml:"regexp,omitempty"`
}
