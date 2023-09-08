package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidEnterpriseWiFiConfigurationWiFiSecurityTypeopen           AndroidEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	AndroidEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise AndroidEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	AndroidEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise  AndroidEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(AndroidEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
	}
}

func parseAndroidEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           AndroidEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wpa2enterprise": AndroidEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpaenterprise":  AndroidEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
