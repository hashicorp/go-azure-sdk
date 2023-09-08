package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationSafeSearchFilter string

const (
	Windows10GeneralConfigurationSafeSearchFiltermoderate    Windows10GeneralConfigurationSafeSearchFilter = "Moderate"
	Windows10GeneralConfigurationSafeSearchFilterstrict      Windows10GeneralConfigurationSafeSearchFilter = "Strict"
	Windows10GeneralConfigurationSafeSearchFilteruserDefined Windows10GeneralConfigurationSafeSearchFilter = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationSafeSearchFilter() []string {
	return []string{
		string(Windows10GeneralConfigurationSafeSearchFiltermoderate),
		string(Windows10GeneralConfigurationSafeSearchFilterstrict),
		string(Windows10GeneralConfigurationSafeSearchFilteruserDefined),
	}
}

func parseWindows10GeneralConfigurationSafeSearchFilter(input string) (*Windows10GeneralConfigurationSafeSearchFilter, error) {
	vals := map[string]Windows10GeneralConfigurationSafeSearchFilter{
		"moderate":    Windows10GeneralConfigurationSafeSearchFiltermoderate,
		"strict":      Windows10GeneralConfigurationSafeSearchFilterstrict,
		"userdefined": Windows10GeneralConfigurationSafeSearchFilteruserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationSafeSearchFilter(input)
	return &out, nil
}
