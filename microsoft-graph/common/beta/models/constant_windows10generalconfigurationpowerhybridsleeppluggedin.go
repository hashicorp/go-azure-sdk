package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerHybridSleepPluggedIn string

const (
	Windows10GeneralConfigurationPowerHybridSleepPluggedIndisabled      Windows10GeneralConfigurationPowerHybridSleepPluggedIn = "Disabled"
	Windows10GeneralConfigurationPowerHybridSleepPluggedInenabled       Windows10GeneralConfigurationPowerHybridSleepPluggedIn = "Enabled"
	Windows10GeneralConfigurationPowerHybridSleepPluggedInnotConfigured Windows10GeneralConfigurationPowerHybridSleepPluggedIn = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationPowerHybridSleepPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerHybridSleepPluggedIndisabled),
		string(Windows10GeneralConfigurationPowerHybridSleepPluggedInenabled),
		string(Windows10GeneralConfigurationPowerHybridSleepPluggedInnotConfigured),
	}
}

func parseWindows10GeneralConfigurationPowerHybridSleepPluggedIn(input string) (*Windows10GeneralConfigurationPowerHybridSleepPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerHybridSleepPluggedIn{
		"disabled":      Windows10GeneralConfigurationPowerHybridSleepPluggedIndisabled,
		"enabled":       Windows10GeneralConfigurationPowerHybridSleepPluggedInenabled,
		"notconfigured": Windows10GeneralConfigurationPowerHybridSleepPluggedInnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerHybridSleepPluggedIn(input)
	return &out, nil
}
