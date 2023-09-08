package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypeopen           AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise  AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
	}
}

func parseAndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wpa2enterprise": AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpaenterprise":  AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
