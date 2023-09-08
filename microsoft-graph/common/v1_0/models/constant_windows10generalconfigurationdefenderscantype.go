package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderScanType string

const (
	Windows10GeneralConfigurationDefenderScanTypedisabled    Windows10GeneralConfigurationDefenderScanType = "Disabled"
	Windows10GeneralConfigurationDefenderScanTypefull        Windows10GeneralConfigurationDefenderScanType = "Full"
	Windows10GeneralConfigurationDefenderScanTypequick       Windows10GeneralConfigurationDefenderScanType = "Quick"
	Windows10GeneralConfigurationDefenderScanTypeuserDefined Windows10GeneralConfigurationDefenderScanType = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderScanType() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderScanTypedisabled),
		string(Windows10GeneralConfigurationDefenderScanTypefull),
		string(Windows10GeneralConfigurationDefenderScanTypequick),
		string(Windows10GeneralConfigurationDefenderScanTypeuserDefined),
	}
}

func parseWindows10GeneralConfigurationDefenderScanType(input string) (*Windows10GeneralConfigurationDefenderScanType, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderScanType{
		"disabled":    Windows10GeneralConfigurationDefenderScanTypedisabled,
		"full":        Windows10GeneralConfigurationDefenderScanTypefull,
		"quick":       Windows10GeneralConfigurationDefenderScanTypequick,
		"userdefined": Windows10GeneralConfigurationDefenderScanTypeuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderScanType(input)
	return &out, nil
}
