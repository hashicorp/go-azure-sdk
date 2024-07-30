package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiConfigurationWifiSecurityType string

const (
	WindowsWifiConfigurationWifiSecurityType_Open           WindowsWifiConfigurationWifiSecurityType = "open"
	WindowsWifiConfigurationWifiSecurityType_Wep            WindowsWifiConfigurationWifiSecurityType = "wep"
	WindowsWifiConfigurationWifiSecurityType_Wpa2Enterprise WindowsWifiConfigurationWifiSecurityType = "wpa2Enterprise"
	WindowsWifiConfigurationWifiSecurityType_Wpa2Personal   WindowsWifiConfigurationWifiSecurityType = "wpa2Personal"
	WindowsWifiConfigurationWifiSecurityType_WpaEnterprise  WindowsWifiConfigurationWifiSecurityType = "wpaEnterprise"
	WindowsWifiConfigurationWifiSecurityType_WpaPersonal    WindowsWifiConfigurationWifiSecurityType = "wpaPersonal"
)

func PossibleValuesForWindowsWifiConfigurationWifiSecurityType() []string {
	return []string{
		string(WindowsWifiConfigurationWifiSecurityType_Open),
		string(WindowsWifiConfigurationWifiSecurityType_Wep),
		string(WindowsWifiConfigurationWifiSecurityType_Wpa2Enterprise),
		string(WindowsWifiConfigurationWifiSecurityType_Wpa2Personal),
		string(WindowsWifiConfigurationWifiSecurityType_WpaEnterprise),
		string(WindowsWifiConfigurationWifiSecurityType_WpaPersonal),
	}
}

func (s *WindowsWifiConfigurationWifiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiConfigurationWifiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiConfigurationWifiSecurityType(input string) (*WindowsWifiConfigurationWifiSecurityType, error) {
	vals := map[string]WindowsWifiConfigurationWifiSecurityType{
		"open":           WindowsWifiConfigurationWifiSecurityType_Open,
		"wep":            WindowsWifiConfigurationWifiSecurityType_Wep,
		"wpa2enterprise": WindowsWifiConfigurationWifiSecurityType_Wpa2Enterprise,
		"wpa2personal":   WindowsWifiConfigurationWifiSecurityType_Wpa2Personal,
		"wpaenterprise":  WindowsWifiConfigurationWifiSecurityType_WpaEnterprise,
		"wpapersonal":    WindowsWifiConfigurationWifiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiConfigurationWifiSecurityType(input)
	return &out, nil
}
