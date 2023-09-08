package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWiFiConfigurationWiFiSecurityType string

const (
	AndroidWiFiConfigurationWiFiSecurityTypeopen           AndroidWiFiConfigurationWiFiSecurityType = "Open"
	AndroidWiFiConfigurationWiFiSecurityTypewpa2Enterprise AndroidWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	AndroidWiFiConfigurationWiFiSecurityTypewpaEnterprise  AndroidWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
)

func PossibleValuesForAndroidWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(AndroidWiFiConfigurationWiFiSecurityTypewpaEnterprise),
	}
}

func parseAndroidWiFiConfigurationWiFiSecurityType(input string) (*AndroidWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidWiFiConfigurationWiFiSecurityType{
		"open":           AndroidWiFiConfigurationWiFiSecurityTypeopen,
		"wpa2enterprise": AndroidWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpaenterprise":  AndroidWiFiConfigurationWiFiSecurityTypewpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
