package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerSleepButtonActionOnBattery string

const (
	Windows10GeneralConfigurationPowerSleepButtonActionOnBatteryhibernate     Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "Hibernate"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBatterynoAction      Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "NoAction"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBatterynotConfigured Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "NotConfigured"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBatteryshutdown      Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "Shutdown"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBatterysleep         Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "Sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerSleepButtonActionOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBatteryhibernate),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBatterynoAction),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBatterynotConfigured),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBatteryshutdown),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBatterysleep),
	}
}

func parseWindows10GeneralConfigurationPowerSleepButtonActionOnBattery(input string) (*Windows10GeneralConfigurationPowerSleepButtonActionOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerSleepButtonActionOnBattery{
		"hibernate":     Windows10GeneralConfigurationPowerSleepButtonActionOnBatteryhibernate,
		"noaction":      Windows10GeneralConfigurationPowerSleepButtonActionOnBatterynoAction,
		"notconfigured": Windows10GeneralConfigurationPowerSleepButtonActionOnBatterynotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerSleepButtonActionOnBatteryshutdown,
		"sleep":         Windows10GeneralConfigurationPowerSleepButtonActionOnBatterysleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerSleepButtonActionOnBattery(input)
	return &out, nil
}
