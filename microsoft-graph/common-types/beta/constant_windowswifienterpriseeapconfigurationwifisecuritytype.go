package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationWifiSecurityType string

const (
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Open           WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "open"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wep            WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "wep"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wpa2Enterprise WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "wpa2Enterprise"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wpa2Personal   WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "wpa2Personal"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_WpaEnterprise  WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "wpaEnterprise"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_WpaPersonal    WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "wpaPersonal"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationWifiSecurityType() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Open),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wep),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wpa2Enterprise),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wpa2Personal),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_WpaEnterprise),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_WpaPersonal),
	}
}

func (s *WindowsWifiEnterpriseEAPConfigurationWifiSecurityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsWifiEnterpriseEAPConfigurationWifiSecurityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsWifiEnterpriseEAPConfigurationWifiSecurityType(input string) (*WindowsWifiEnterpriseEAPConfigurationWifiSecurityType, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationWifiSecurityType{
		"open":           WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Open,
		"wep":            WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wep,
		"wpa2enterprise": WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wpa2Enterprise,
		"wpa2personal":   WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_Wpa2Personal,
		"wpaenterprise":  WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_WpaEnterprise,
		"wpapersonal":    WindowsWifiEnterpriseEAPConfigurationWifiSecurityType_WpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationWifiSecurityType(input)
	return &out, nil
}
