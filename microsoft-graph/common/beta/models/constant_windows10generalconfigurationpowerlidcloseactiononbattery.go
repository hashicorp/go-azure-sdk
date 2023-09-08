package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerLidCloseActionOnBattery string

const (
	Windows10GeneralConfigurationPowerLidCloseActionOnBatteryhibernate     Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "Hibernate"
	Windows10GeneralConfigurationPowerLidCloseActionOnBatterynoAction      Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "NoAction"
	Windows10GeneralConfigurationPowerLidCloseActionOnBatterynotConfigured Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "NotConfigured"
	Windows10GeneralConfigurationPowerLidCloseActionOnBatteryshutdown      Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "Shutdown"
	Windows10GeneralConfigurationPowerLidCloseActionOnBatterysleep         Windows10GeneralConfigurationPowerLidCloseActionOnBattery = "Sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerLidCloseActionOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBatteryhibernate),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBatterynoAction),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBatterynotConfigured),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBatteryshutdown),
		string(Windows10GeneralConfigurationPowerLidCloseActionOnBatterysleep),
	}
}

func parseWindows10GeneralConfigurationPowerLidCloseActionOnBattery(input string) (*Windows10GeneralConfigurationPowerLidCloseActionOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerLidCloseActionOnBattery{
		"hibernate":     Windows10GeneralConfigurationPowerLidCloseActionOnBatteryhibernate,
		"noaction":      Windows10GeneralConfigurationPowerLidCloseActionOnBatterynoAction,
		"notconfigured": Windows10GeneralConfigurationPowerLidCloseActionOnBatterynotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerLidCloseActionOnBatteryshutdown,
		"sleep":         Windows10GeneralConfigurationPowerLidCloseActionOnBatterysleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerLidCloseActionOnBattery(input)
	return &out, nil
}
