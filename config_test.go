//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config_test

import (
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/tdrn-org/go-config-toml"
)

func TestDefaults(t *testing.T) {
	cfg := &Config{}
	err := config.Defaults(defaultsData, cfg)
	require.NoError(t, err)
}

func TestSave(t *testing.T) {
	cfg1 := &Config{}
	err := config.Load(cfg1, "testdata/config_valid.toml", defaultsData, true)
	require.NoError(t, err)
	cfgPath := filepath.Join(t.TempDir(), "config.toml")
	err = config.Save(cfg1, cfgPath, 0600)
	require.NoError(t, err)
	cfg2 := &Config{}
	err = config.Load(cfg2, cfgPath, defaultsData, true)
	require.NoError(t, err)
	require.Equal(t, cfg1, cfg2)
}

func TestLoadValid(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_valid.toml", defaultsData, true)
	require.NoError(t, err)
}

func TestLoadInvalidDuration(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_duration.toml", defaultsData, true)
	require.Error(t, err)
}

func TestLoadInvalidNetworks(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_networks.toml", defaultsData, true)
	require.Error(t, err)
}

func TestLoadInvalidRegexp(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_regexp.toml", defaultsData, true)
	require.Error(t, err)
}

func TestLoadInvalidTimeLocation(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_time_location.toml", defaultsData, true)
	require.Error(t, err)
}

func TestLoadInvalidURLs(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_urls.toml", defaultsData, true)
	require.Error(t, err)
}

func TestLoadInvalidEnum(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_enum.toml", defaultsData, true)
	require.Error(t, err)
}

func TestLoadInvalidUnknownField(t *testing.T) {
	cfg := &Config{}
	err := config.Load(cfg, "testdata/config_invalid_unknown_field.toml", defaultsData, true)
	require.Error(t, err)
}

var defaultsData []byte = []byte(`
duration = ""
networks = [ ]
regexp = ""
time_location = ""
urls = [ "" ]
enum = "alpha"
`)

type Config struct {
	Duration     config.DurationSpec     `toml:"duration,omitempty"`
	Networks     config.NetworkSpecs     `toml:"networks,omitempty"`
	TimeLocation config.TimeLocationSpec `toml:"time_location,omitempty"`
	URLs         config.URLSpecs         `toml:"urls,omitempty"`
	Regexp       config.RegexpSpec       `toml:"regexp,omitempty"`
	Enum         MyEnum                  `toml:"enum"`
}

type MyEnum int

const (
	MyEnumAlpha MyEnum = iota
	MyEnumBeta
)

var myEnumMarshalMap map[MyEnum]string = map[MyEnum]string{
	MyEnumAlpha: "alpha",
	MyEnumBeta:  "beta",
}

var myEnumUnmarshalMap map[string]MyEnum = map[string]MyEnum{
	"alpha": MyEnumAlpha,
	"beta":  MyEnumBeta,
}

func (e MyEnum) MarshalText() ([]byte, error) {
	return config.MarshalEnum(e, myEnumMarshalMap)
}

func (e *MyEnum) UnmarshalText(text []byte) error {
	var err error
	*e, err = config.UnmarshalEnum(myEnumUnmarshalMap, text)
	return err
}
