package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationTunnelConfiguration string

const (
	AppleVpnAlwaysOnConfigurationTunnelConfigurationcellular        AppleVpnAlwaysOnConfigurationTunnelConfiguration = "Cellular"
	AppleVpnAlwaysOnConfigurationTunnelConfigurationwifi            AppleVpnAlwaysOnConfigurationTunnelConfiguration = "Wifi"
	AppleVpnAlwaysOnConfigurationTunnelConfigurationwifiAndCellular AppleVpnAlwaysOnConfigurationTunnelConfiguration = "WifiAndCellular"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationTunnelConfiguration() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationTunnelConfigurationcellular),
		string(AppleVpnAlwaysOnConfigurationTunnelConfigurationwifi),
		string(AppleVpnAlwaysOnConfigurationTunnelConfigurationwifiAndCellular),
	}
}

func parseAppleVpnAlwaysOnConfigurationTunnelConfiguration(input string) (*AppleVpnAlwaysOnConfigurationTunnelConfiguration, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationTunnelConfiguration{
		"cellular":        AppleVpnAlwaysOnConfigurationTunnelConfigurationcellular,
		"wifi":            AppleVpnAlwaysOnConfigurationTunnelConfigurationwifi,
		"wifiandcellular": AppleVpnAlwaysOnConfigurationTunnelConfigurationwifiAndCellular,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationTunnelConfiguration(input)
	return &out, nil
}
