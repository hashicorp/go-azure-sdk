package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiFiConfigurationWiFiSecurityType string

const (
	MacOSWiFiConfigurationWiFiSecurityTypeopen           MacOSWiFiConfigurationWiFiSecurityType = "Open"
	MacOSWiFiConfigurationWiFiSecurityTypewep            MacOSWiFiConfigurationWiFiSecurityType = "Wep"
	MacOSWiFiConfigurationWiFiSecurityTypewpa2Enterprise MacOSWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	MacOSWiFiConfigurationWiFiSecurityTypewpa2Personal   MacOSWiFiConfigurationWiFiSecurityType = "Wpa2Personal"
	MacOSWiFiConfigurationWiFiSecurityTypewpaEnterprise  MacOSWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	MacOSWiFiConfigurationWiFiSecurityTypewpaPersonal    MacOSWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForMacOSWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(MacOSWiFiConfigurationWiFiSecurityTypeopen),
		string(MacOSWiFiConfigurationWiFiSecurityTypewep),
		string(MacOSWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(MacOSWiFiConfigurationWiFiSecurityTypewpa2Personal),
		string(MacOSWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(MacOSWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseMacOSWiFiConfigurationWiFiSecurityType(input string) (*MacOSWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]MacOSWiFiConfigurationWiFiSecurityType{
		"open":           MacOSWiFiConfigurationWiFiSecurityTypeopen,
		"wep":            MacOSWiFiConfigurationWiFiSecurityTypewep,
		"wpa2enterprise": MacOSWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpa2personal":   MacOSWiFiConfigurationWiFiSecurityTypewpa2Personal,
		"wpaenterprise":  MacOSWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":    MacOSWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
