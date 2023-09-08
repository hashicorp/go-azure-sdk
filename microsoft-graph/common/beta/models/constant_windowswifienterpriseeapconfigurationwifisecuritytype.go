package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationWifiSecurityType string

const (
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypeopen           WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "Open"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewep            WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "Wep"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpa2Enterprise WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "Wpa2Enterprise"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpa2Personal   WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "Wpa2Personal"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpaEnterprise  WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "WpaEnterprise"
	WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpaPersonal    WindowsWifiEnterpriseEAPConfigurationWifiSecurityType = "WpaPersonal"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationWifiSecurityType() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypeopen),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewep),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpa2Enterprise),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpa2Personal),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpaEnterprise),
		string(WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpaPersonal),
	}
}

func parseWindowsWifiEnterpriseEAPConfigurationWifiSecurityType(input string) (*WindowsWifiEnterpriseEAPConfigurationWifiSecurityType, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationWifiSecurityType{
		"open":           WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypeopen,
		"wep":            WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewep,
		"wpa2enterprise": WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpa2Enterprise,
		"wpa2personal":   WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpa2Personal,
		"wpaenterprise":  WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpaEnterprise,
		"wpapersonal":    WindowsWifiEnterpriseEAPConfigurationWifiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationWifiSecurityType(input)
	return &out, nil
}
