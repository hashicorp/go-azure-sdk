package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	IosEnterpriseWiFiConfigurationWiFiSecurityTypeopen           IosEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	IosEnterpriseWiFiConfigurationWiFiSecurityTypewep            IosEnterpriseWiFiConfigurationWiFiSecurityType = "Wep"
	IosEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise IosEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	IosEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Personal   IosEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Personal"
	IosEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise  IosEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
	IosEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal    IosEnterpriseWiFiConfigurationWiFiSecurityType = "WpaPersonal"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityTypewep),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Personal),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
		string(IosEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal),
	}
}

func parseIosEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*IosEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           IosEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wep":            IosEnterpriseWiFiConfigurationWiFiSecurityTypewep,
		"wpa2enterprise": IosEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpa2personal":   IosEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Personal,
		"wpaenterprise":  IosEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
		"wpapersonal":    IosEnterpriseWiFiConfigurationWiFiSecurityTypewpaPersonal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
