package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EndpointProtectionConfigurationDefenderScanType string

const (
	Windows10EndpointProtectionConfigurationDefenderScanTypedisabled    Windows10EndpointProtectionConfigurationDefenderScanType = "Disabled"
	Windows10EndpointProtectionConfigurationDefenderScanTypefull        Windows10EndpointProtectionConfigurationDefenderScanType = "Full"
	Windows10EndpointProtectionConfigurationDefenderScanTypequick       Windows10EndpointProtectionConfigurationDefenderScanType = "Quick"
	Windows10EndpointProtectionConfigurationDefenderScanTypeuserDefined Windows10EndpointProtectionConfigurationDefenderScanType = "UserDefined"
)

func PossibleValuesForWindows10EndpointProtectionConfigurationDefenderScanType() []string {
	return []string{
		string(Windows10EndpointProtectionConfigurationDefenderScanTypedisabled),
		string(Windows10EndpointProtectionConfigurationDefenderScanTypefull),
		string(Windows10EndpointProtectionConfigurationDefenderScanTypequick),
		string(Windows10EndpointProtectionConfigurationDefenderScanTypeuserDefined),
	}
}

func parseWindows10EndpointProtectionConfigurationDefenderScanType(input string) (*Windows10EndpointProtectionConfigurationDefenderScanType, error) {
	vals := map[string]Windows10EndpointProtectionConfigurationDefenderScanType{
		"disabled":    Windows10EndpointProtectionConfigurationDefenderScanTypedisabled,
		"full":        Windows10EndpointProtectionConfigurationDefenderScanTypefull,
		"quick":       Windows10EndpointProtectionConfigurationDefenderScanTypequick,
		"userdefined": Windows10EndpointProtectionConfigurationDefenderScanTypeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EndpointProtectionConfigurationDefenderScanType(input)
	return &out, nil
}
