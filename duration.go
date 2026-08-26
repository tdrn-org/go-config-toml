//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
	"time"
)

// DurationSpec supports Marshaling of [time.Duration] values.
type DurationSpec time.Duration

// See [encoding.TextMarshaler]
func (spec *DurationSpec) MarshalText() ([]byte, error) {
	return []byte(time.Duration(*spec).String()), nil
}

// See [encoding.TextUnmarshaler]
func (spec *DurationSpec) UnmarshalText(text []byte) error {
	durationString := string(text)
	if durationString == "" {
		*spec = DurationSpec(0)
		return nil
	}
	parsedDuration, err := time.ParseDuration(durationString)
	if err != nil {
		return fmt.Errorf("invalid duration '%s' (cause: %w)", durationString, err)
	}
	*spec = DurationSpec(parsedDuration)
	return nil
}
