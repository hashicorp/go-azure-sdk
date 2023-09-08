package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosGeneralDeviceConfigurationMediaContentRatingApps string

const (
	IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove12 IosGeneralDeviceConfigurationMediaContentRatingApps = "AgesAbove12"
	IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove17 IosGeneralDeviceConfigurationMediaContentRatingApps = "AgesAbove17"
	IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove4  IosGeneralDeviceConfigurationMediaContentRatingApps = "AgesAbove4"
	IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove9  IosGeneralDeviceConfigurationMediaContentRatingApps = "AgesAbove9"
	IosGeneralDeviceConfigurationMediaContentRatingAppsallAllowed  IosGeneralDeviceConfigurationMediaContentRatingApps = "AllAllowed"
	IosGeneralDeviceConfigurationMediaContentRatingAppsallBlocked  IosGeneralDeviceConfigurationMediaContentRatingApps = "AllBlocked"
)

func PossibleValuesForIosGeneralDeviceConfigurationMediaContentRatingApps() []string {
	return []string{
		string(IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove12),
		string(IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove17),
		string(IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove4),
		string(IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove9),
		string(IosGeneralDeviceConfigurationMediaContentRatingAppsallAllowed),
		string(IosGeneralDeviceConfigurationMediaContentRatingAppsallBlocked),
	}
}

func parseIosGeneralDeviceConfigurationMediaContentRatingApps(input string) (*IosGeneralDeviceConfigurationMediaContentRatingApps, error) {
	vals := map[string]IosGeneralDeviceConfigurationMediaContentRatingApps{
		"agesabove12": IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove12,
		"agesabove17": IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove17,
		"agesabove4":  IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove4,
		"agesabove9":  IosGeneralDeviceConfigurationMediaContentRatingAppsagesAbove9,
		"allallowed":  IosGeneralDeviceConfigurationMediaContentRatingAppsallAllowed,
		"allblocked":  IosGeneralDeviceConfigurationMediaContentRatingAppsallBlocked,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosGeneralDeviceConfigurationMediaContentRatingApps(input)
	return &out, nil
}
