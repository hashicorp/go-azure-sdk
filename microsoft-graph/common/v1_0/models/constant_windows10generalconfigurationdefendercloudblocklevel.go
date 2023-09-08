package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderCloudBlockLevel string

const (
	Windows10GeneralConfigurationDefenderCloudBlockLevelhigh          Windows10GeneralConfigurationDefenderCloudBlockLevel = "High"
	Windows10GeneralConfigurationDefenderCloudBlockLevelhighPlus      Windows10GeneralConfigurationDefenderCloudBlockLevel = "HighPlus"
	Windows10GeneralConfigurationDefenderCloudBlockLevelnotConfigured Windows10GeneralConfigurationDefenderCloudBlockLevel = "NotConfigured"
	Windows10GeneralConfigurationDefenderCloudBlockLevelzeroTolerance Windows10GeneralConfigurationDefenderCloudBlockLevel = "ZeroTolerance"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderCloudBlockLevel() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderCloudBlockLevelhigh),
		string(Windows10GeneralConfigurationDefenderCloudBlockLevelhighPlus),
		string(Windows10GeneralConfigurationDefenderCloudBlockLevelnotConfigured),
		string(Windows10GeneralConfigurationDefenderCloudBlockLevelzeroTolerance),
	}
}

func parseWindows10GeneralConfigurationDefenderCloudBlockLevel(input string) (*Windows10GeneralConfigurationDefenderCloudBlockLevel, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderCloudBlockLevel{
		"high":          Windows10GeneralConfigurationDefenderCloudBlockLevelhigh,
		"highplus":      Windows10GeneralConfigurationDefenderCloudBlockLevelhighPlus,
		"notconfigured": Windows10GeneralConfigurationDefenderCloudBlockLevelnotConfigured,
		"zerotolerance": Windows10GeneralConfigurationDefenderCloudBlockLevelzeroTolerance,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderCloudBlockLevel(input)
	return &out, nil
}
