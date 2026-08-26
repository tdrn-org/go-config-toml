//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/pelletier/go-toml/v2"
)

func Defaults(defaultsData []byte, cfg any) error {
	err := toml.Unmarshal(defaultsData, cfg)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config defaults (cause: %w)", err)
	}
	return nil
}

func Load(path string, defaultsData []byte, cfg any, strict bool) error {
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
