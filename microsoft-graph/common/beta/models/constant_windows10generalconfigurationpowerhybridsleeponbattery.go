package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerHybridSleepOnBattery string

const (
	Windows10GeneralConfigurationPowerHybridSleepOnBatterydisabled      Windows10GeneralConfigurationPowerHybridSleepOnBattery = "Disabled"
	Windows10GeneralConfigurationPowerHybridSleepOnBatteryenabled       Windows10GeneralConfigurationPowerHybridSleepOnBattery = "Enabled"
	Windows10GeneralConfigurationPowerHybridSleepOnBatterynotConfigured Windows10GeneralConfigurationPowerHybridSleepOnBattery = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationPowerHybridSleepOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerHybridSleepOnBatterydisabled),
		string(Windows10GeneralConfigurationPowerHybridSleepOnBatteryenabled),
		string(Windows10GeneralConfigurationPowerHybridSleepOnBatterynotConfigured),
	}
}

func parseWindows10GeneralConfigurationPowerHybridSleepOnBattery(input string) (*Windows10GeneralConfigurationPowerHybridSleepOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerHybridSleepOnBattery{
		"disabled":      Windows10GeneralConfigurationPowerHybridSleepOnBatterydisabled,
		"enabled":       Windows10GeneralConfigurationPowerHybridSleepOnBatteryenabled,
		"notconfigured": Windows10GeneralConfigurationPowerHybridSleepOnBatterynotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerHybridSleepOnBattery(input)
	return &out, nil
}
