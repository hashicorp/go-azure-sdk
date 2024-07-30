package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsDeliveryOptimizationConfigurationVpnPeerCaching string

const (
	WindowsDeliveryOptimizationConfigurationVpnPeerCaching_Disabled      WindowsDeliveryOptimizationConfigurationVpnPeerCaching = "disabled"
	WindowsDeliveryOptimizationConfigurationVpnPeerCaching_Enabled       WindowsDeliveryOptimizationConfigurationVpnPeerCaching = "enabled"
	WindowsDeliveryOptimizationConfigurationVpnPeerCaching_NotConfigured WindowsDeliveryOptimizationConfigurationVpnPeerCaching = "notConfigured"
)

func PossibleValuesForWindowsDeliveryOptimizationConfigurationVpnPeerCaching() []string {
	return []string{
		string(WindowsDeliveryOptimizationConfigurationVpnPeerCaching_Disabled),
		string(WindowsDeliveryOptimizationConfigurationVpnPeerCaching_Enabled),
		string(WindowsDeliveryOptimizationConfigurationVpnPeerCaching_NotConfigured),
	}
}

func (s *WindowsDeliveryOptimizationConfigurationVpnPeerCaching) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsDeliveryOptimizationConfigurationVpnPeerCaching(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsDeliveryOptimizationConfigurationVpnPeerCaching(input string) (*WindowsDeliveryOptimizationConfigurationVpnPeerCaching, error) {
	vals := map[string]WindowsDeliveryOptimizationConfigurationVpnPeerCaching{
		"disabled":      WindowsDeliveryOptimizationConfigurationVpnPeerCaching_Disabled,
		"enabled":       WindowsDeliveryOptimizationConfigurationVpnPeerCaching_Enabled,
		"notconfigured": WindowsDeliveryOptimizationConfigurationVpnPeerCaching_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsDeliveryOptimizationConfigurationVpnPeerCaching(input)
	return &out, nil
}
