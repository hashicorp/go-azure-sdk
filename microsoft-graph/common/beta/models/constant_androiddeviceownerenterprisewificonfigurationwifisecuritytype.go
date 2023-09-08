package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypeopen          AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewep           AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "Wep"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal   AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewep),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":          AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wep":           AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewep,
		"wpaenterprise": AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":   AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
