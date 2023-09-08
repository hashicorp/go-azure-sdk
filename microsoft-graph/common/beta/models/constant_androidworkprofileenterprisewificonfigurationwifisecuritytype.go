package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypeopen           AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType = "Open"
	AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise  AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise),
	}
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType{
		"open":           AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypeopen,
		"wpa2enterprise": AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpaenterprise":  AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityTypewpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
