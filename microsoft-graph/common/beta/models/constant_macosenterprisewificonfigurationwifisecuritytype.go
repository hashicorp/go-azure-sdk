package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	MacOSEnterpriseWiFiConfigurationWiFiSecurityTypeopen           MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewep            MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "Wep"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Personal   MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Personal"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise  MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal    MacOSEnterpriseWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewep),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Personal),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseMacOSEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*MacOSEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           MacOSEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wep":            MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewep,
		"wpa2enterprise": MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpa2personal":   MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Personal,
		"wpaenterprise":  MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":    MacOSEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
