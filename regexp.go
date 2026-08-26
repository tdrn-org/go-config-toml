//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
	"regexp"
)

// RegexpSpec supports Marshaling of [regexp.Regexp] values.
type RegexpSpec struct {
	*regexp.Regexp
}

// See [encoding.TextMarshaler]
func (spec *RegexpSpec) MarshalText() ([]byte, error) {
	if spec.Regexp == nil {
		return []byte(""), nil
	}
	return []byte(spec.Regexp.String()), nil
}

// See [encoding.TextUnmarshaler]
func (spec *RegexpSpec) UnmarshalText(text []byte) error {
	regexpString := string(text)
	if regexpString == "" {
		spec.Regexp = nil
		return nil
	}
	parsedRegexp, err := regexp.Compile(regexpString)
	if err != nil {
		return fmt.Errorf("invalid Regexp: '%s' (cause: %w)", regexpString, err)
	}
	spec.Regexp = parsedRegexp
	return nil
}
