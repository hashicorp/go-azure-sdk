package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypeopen          AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewep           AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "Wep"
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal   AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewep),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":          AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wep":           AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewep,
		"wpaenterprise": AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":   AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
