//
// Copyright (C) 2026 Holger de Carne
//
// This software may be modified and distributed under the terms
// of the MIT license. See the LICENSE file for details.

package config

import (
	"fmt"
	"net/netip"
)

type NetworkSpec struct {
	netip.Prefix
}

func (spec *NetworkSpec) MarshalText() ([]byte, error) {
	return []byte(spec.Prefix.String()), nil
}

func (spec *NetworkSpec) UnmarshalText(text []byte) error {
	networkString := string(text)
	parsedNetwork, err := netip.ParsePrefix(networkString)
	if err != nil {
		return fmt.Errorf("invalid network: '%s' (cause: %w)", networkString, err)
	}
	spec.Prefix = parsedNetwork
	return nil
}

type NetworkSpecs []NetworkSpec

func (specs NetworkSpecs) Prefixes() []netip.Prefix {
	networks := make([]netip.Prefix, 0, len(specs))
	for _, spec := range specs {
		networks = append(networks, spec.Prefix)
	}
	return networks
}
