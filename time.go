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

type TimeLocationSpec struct {
	*time.Location
}

func (spec *TimeLocationSpec) MarshalText() ([]byte, error) {
	if spec.Location == nil {
		return []byte(""), nil
	}
	return []byte(spec.Location.String()), nil
}

func (spec *TimeLocationSpec) UnmarshalText(text []byte) error {
	locationString := string(text)
	if locationString == "" {
		spec.Location = nil
		return nil
	}
	parsedLocation, err := time.LoadLocation(locationString)
	if err != nil {
		return fmt.Errorf("invalid time location: '%s' (cause: %w)", locationString, err)
	}
	spec.Location = parsedLocation
	return nil
}
