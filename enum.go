//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
)

// MarshalEnum is a helper function to implement [encoding.TextMarshaler]
// for an enum type based on the given map.
func MarshalEnum[E comparable](e E, marshalMap map[E]string) ([]byte, error) {
	s, ok := marshalMap[e]
	if !ok {
		return nil, fmt.Errorf("undefined enum: '%v'", e)
	}
	return []byte(s), nil
}

// MarshalStringerEnum is a helper function to implement [encoding.TextMarshaler]
// for an enum type based on the enum type's [fmt.Stringer] interface.
func MarshalStringerEnum[E fmt.Stringer](e E) ([]byte, error) {
	return []byte(e.String()), nil
}

// UnmarshalEnum is a helper function to implement [encoding.TextUnmarshaler]
// for an enum type.
func UnmarshalEnum[E comparable](unmarshalMap map[string]E, text []byte) (E, error) {
	var defaultEnum E
	enumString := string(text)
	if enumString == "" {
		return defaultEnum, nil
	}
	unmarshaledEnum, ok := unmarshalMap[enumString]
	if !ok {
		return defaultEnum, fmt.Errorf("unknown enum: '%s'", enumString)
	}
	return unmarshaledEnum, nil
}
