package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerWiFiConfigurationWiFiSecurityType string

const (
	AospDeviceOwnerWiFiConfigurationWiFiSecurityTypeopen          AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "Open"
	AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewep           AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "Wep"
	AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaEnterprise AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaPersonal   AospDeviceOwnerWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForAospDeviceOwnerWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityTypeopen),
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewep),
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseAospDeviceOwnerWiFiConfigurationWiFiSecurityType(input string) (*AospDeviceOwnerWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AospDeviceOwnerWiFiConfigurationWiFiSecurityType{
		"open":          AospDeviceOwnerWiFiConfigurationWiFiSecurityTypeopen,
		"wep":           AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewep,
		"wpaenterprise": AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":   AospDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
