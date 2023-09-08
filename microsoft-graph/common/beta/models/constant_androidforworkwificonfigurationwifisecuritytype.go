package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkWiFiConfigurationWiFiSecurityType string

const (
	AndroidForWorkWiFiConfigurationWiFiSecurityTypeopen           AndroidForWorkWiFiConfigurationWiFiSecurityType = "Open"
	AndroidForWorkWiFiConfigurationWiFiSecurityTypewpa2Enterprise AndroidForWorkWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	AndroidForWorkWiFiConfigurationWiFiSecurityTypewpaEnterprise  AndroidForWorkWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
)

func PossibleValuesForAndroidForWorkWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidForWorkWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidForWorkWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(AndroidForWorkWiFiConfigurationWiFiSecurityTypewpaEnterprise),
	}
}

func parseAndroidForWorkWiFiConfigurationWiFiSecurityType(input string) (*AndroidForWorkWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidForWorkWiFiConfigurationWiFiSecurityType{
		"open":           AndroidForWorkWiFiConfigurationWiFiSecurityTypeopen,
		"wpa2enterprise": AndroidForWorkWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpaenterprise":  AndroidForWorkWiFiConfigurationWiFiSecurityTypewpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
