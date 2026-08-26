//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
)

// MarshalEnum is a helper function to implement [encoding.TestMarshaler]
// for enums.
func MarshalEnum[E comparable](e E, marshalMap map[E]string) ([]byte, error) {
	s, ok := marshalMap[e]
	if !ok {
		return []byte(""), nil
	}
	return []byte(s), nil
}

// MarshalEnum is a helper function to implement [encoding.TestUnmarshaler]
// for enums.
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
