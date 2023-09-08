package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileWiFiConfigurationWiFiSecurityType string

const (
	AndroidWorkProfileWiFiConfigurationWiFiSecurityTypeopen           AndroidWorkProfileWiFiConfigurationWiFiSecurityType = "Open"
	AndroidWorkProfileWiFiConfigurationWiFiSecurityTypewpa2Enterprise AndroidWorkProfileWiFiConfigurationWiFiSecurityType = "Wpa2Enterprise"
	AndroidWorkProfileWiFiConfigurationWiFiSecurityTypewpaEnterprise  AndroidWorkProfileWiFiConfigurationWiFiSecurityType = "WpaEnterprise"
)

func PossibleValuesForAndroidWorkProfileWiFiConfigurationWiFiSecurityType() []string {
	return []string{
		string(AndroidWorkProfileWiFiConfigurationWiFiSecurityTypeopen),
		string(AndroidWorkProfileWiFiConfigurationWiFiSecurityTypewpa2Enterprise),
		string(AndroidWorkProfileWiFiConfigurationWiFiSecurityTypewpaEnterprise),
	}
}

func parseAndroidWorkProfileWiFiConfigurationWiFiSecurityType(input string) (*AndroidWorkProfileWiFiConfigurationWiFiSecurityType, error) {
	vals := map[string]AndroidWorkProfileWiFiConfigurationWiFiSecurityType{
		"open":           AndroidWorkProfileWiFiConfigurationWiFiSecurityTypeopen,
		"wpa2enterprise": AndroidWorkProfileWiFiConfigurationWiFiSecurityTypewpa2Enterprise,
		"wpaenterprise":  AndroidWorkProfileWiFiConfigurationWiFiSecurityTypewpaEnterprise,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileWiFiConfigurationWiFiSecurityType(input)
	return &out, nil
}
