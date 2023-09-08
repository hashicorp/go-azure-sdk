package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn string

const (
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedInhibernate     Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "Hibernate"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedInnoAction      Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "NoAction"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedInnotConfigured Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "NotConfigured"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedInshutdown      Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "Shutdown"
	Windows10GeneralConfigurationPowerSleepButtonActionPluggedInsleep         Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn = "Sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerSleepButtonActionPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedInhibernate),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedInnoAction),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedInnotConfigured),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedInshutdown),
		string(Windows10GeneralConfigurationPowerSleepButtonActionPluggedInsleep),
	}
}

func parseWindows10GeneralConfigurationPowerSleepButtonActionPluggedIn(input string) (*Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn{
		"hibernate":     Windows10GeneralConfigurationPowerSleepButtonActionPluggedInhibernate,
		"noaction":      Windows10GeneralConfigurationPowerSleepButtonActionPluggedInnoAction,
		"notconfigured": Windows10GeneralConfigurationPowerSleepButtonActionPluggedInnotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerSleepButtonActionPluggedInshutdown,
		"sleep":         Windows10GeneralConfigurationPowerSleepButtonActionPluggedInsleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerSleepButtonActionPluggedIn(input)
	return &out, nil
}
