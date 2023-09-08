package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType string

const (
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypeopen          AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "Open"
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewep           AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "Wep"
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaEnterprise AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaPersonal   AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForAndroidDeviceOwnerWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewep),
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseAndroidDeviceOwnerWiFiConfigurationWiFiSecurityType(input string) (*AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType{
		"open":          AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypeopen,
		"wep":           AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewep,
		"wpaenterprise": AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":   AndroidDeviceOwnerWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
