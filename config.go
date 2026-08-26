//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

// Package config provides functions and types to define config
// objects and load and save them to TOML files.
package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pelletier/go-toml/v2"
)

// Defaults creates a default config from the given
// default config data (usually statically embedded).
func Defaults(defaultsData []byte, cfg any) error {
	err := toml.Unmarshal(defaultsData, cfg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config defaults (cause: %w)", err)
	}
	return nil
}

// Load loads a config from a file. Before loading
// the config from the given path, the config object
// is initialized using the given default data.
//
// If the strict flag is set non-matching fields
// will cause errors. Otherwise they are silently
// ignored.
func Load(cfg any, path string, defaultsData []byte, strict bool) error {
	logger := slog.With(slog.String("path", path))
	logger.Info("loading config")
	err := Defaults(defaultsData, cfg)
	if err != nil {
		return err
	}
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("failed to open config file '%s' (cause: %w)", path, err)
	}
	defer file.Close()
	decoder := toml.NewDecoder(file)
	if strict {
		decoder = decoder.DisallowUnknownFields()
	}
	err = decoder.Decode(cfg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config file '%s' (cause: %w)", path, err)
	}
	return nil
}

// Save saves the given config object to the given path.
func Save(cfg any, path string, perm os.FileMode) error {
	logger := slog.With(slog.String("path", path))
	logger.Info("saving config")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, perm)
	if err != nil {
		return fmt.Errorf("failed to create/truncate config file '%s' (cause: %w)", path, err)
	}
	defer file.Close()
	err = toml.NewEncoder(file).Encode(cfg)
	if err != nil {
		return fmt.Errorf("failed to marshal config to file '%s' (cause: %w)", path, err)
	}
	return nil
}
