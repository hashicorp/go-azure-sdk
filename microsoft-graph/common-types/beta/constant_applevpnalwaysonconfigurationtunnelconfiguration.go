package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnAlwaysOnConfigurationTunnelConfiguration string

const (
	AppleVpnAlwaysOnConfigurationTunnelConfiguration_Cellular        AppleVpnAlwaysOnConfigurationTunnelConfiguration = "cellular"
	AppleVpnAlwaysOnConfigurationTunnelConfiguration_Wifi            AppleVpnAlwaysOnConfigurationTunnelConfiguration = "wifi"
	AppleVpnAlwaysOnConfigurationTunnelConfiguration_WifiAndCellular AppleVpnAlwaysOnConfigurationTunnelConfiguration = "wifiAndCellular"
)

func PossibleValuesForAppleVpnAlwaysOnConfigurationTunnelConfiguration() []string {
	return []string{
		string(AppleVpnAlwaysOnConfigurationTunnelConfiguration_Cellular),
		string(AppleVpnAlwaysOnConfigurationTunnelConfiguration_Wifi),
		string(AppleVpnAlwaysOnConfigurationTunnelConfiguration_WifiAndCellular),
	}
}

func (s *AppleVpnAlwaysOnConfigurationTunnelConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleVpnAlwaysOnConfigurationTunnelConfiguration(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleVpnAlwaysOnConfigurationTunnelConfiguration(input string) (*AppleVpnAlwaysOnConfigurationTunnelConfiguration, error) {
	vals := map[string]AppleVpnAlwaysOnConfigurationTunnelConfiguration{
		"cellular":        AppleVpnAlwaysOnConfigurationTunnelConfiguration_Cellular,
		"wifi":            AppleVpnAlwaysOnConfigurationTunnelConfiguration_Wifi,
		"wifiandcellular": AppleVpnAlwaysOnConfigurationTunnelConfiguration_WifiAndCellular,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnAlwaysOnConfigurationTunnelConfiguration(input)
	return &out, nil
}
