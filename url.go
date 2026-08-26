//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
	"net/url"
)

type URLSpec struct {
	*url.URL
}

func (spec *URLSpec) MarshalText() ([]byte, error) {
	if spec.URL == nil {
		return []byte(""), nil
	}
	return []byte(spec.URL.String()), nil
}

func (spec *URLSpec) UnmarshalText(text []byte) error {
	urlString := string(text)
	if urlString == "" {
		spec.URL = nil
		return nil
	}
	parsedURL, err := url.Parse(urlString)
	if err != nil {
		return fmt.Errorf("invalid URL: '%s' (cause: %w)", urlString, err)
	}
	spec.URL = parsedURL
	return nil
}

type URLSpecs []URLSpec

func (specs URLSpecs) URLs() []*url.URL {
	urls := make([]*url.URL, 0, len(specs))
	for _, spec := range specs {
		urls = append(urls, spec.URL)
	}
	return urls
}
