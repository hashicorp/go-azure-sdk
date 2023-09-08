package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration string

const (
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationbottom        AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "Bottom"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationhide          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "Hide"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationnotConfigured AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "NotConfigured"
	AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationtop           AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration = "Top"
)

func PossibleValuesForAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration() []string {
	return []string{
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationbottom),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationhide),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationnotConfigured),
		string(AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationtop),
	}
}

func parseAndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration(input string) (*AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration, error) {
	vals := map[string]AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration{
		"bottom":        AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationbottom,
		"hide":          AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationhide,
		"notconfigured": AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationnotConfigured,
		"top":           AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfigurationtop,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerGeneralDeviceConfigurationMicrosoftLauncherSearchBarPlacementConfiguration(input)
	return &out, nil
}
