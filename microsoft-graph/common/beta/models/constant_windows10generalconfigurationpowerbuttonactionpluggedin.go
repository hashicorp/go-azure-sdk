package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerButtonActionPluggedIn string

const (
	Windows10GeneralConfigurationPowerButtonActionPluggedInhibernate     Windows10GeneralConfigurationPowerButtonActionPluggedIn = "Hibernate"
	Windows10GeneralConfigurationPowerButtonActionPluggedInnoAction      Windows10GeneralConfigurationPowerButtonActionPluggedIn = "NoAction"
	Windows10GeneralConfigurationPowerButtonActionPluggedInnotConfigured Windows10GeneralConfigurationPowerButtonActionPluggedIn = "NotConfigured"
	Windows10GeneralConfigurationPowerButtonActionPluggedInshutdown      Windows10GeneralConfigurationPowerButtonActionPluggedIn = "Shutdown"
	Windows10GeneralConfigurationPowerButtonActionPluggedInsleep         Windows10GeneralConfigurationPowerButtonActionPluggedIn = "Sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerButtonActionPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerButtonActionPluggedInhibernate),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedInnoAction),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedInnotConfigured),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedInshutdown),
		string(Windows10GeneralConfigurationPowerButtonActionPluggedInsleep),
	}
}

func parseWindows10GeneralConfigurationPowerButtonActionPluggedIn(input string) (*Windows10GeneralConfigurationPowerButtonActionPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerButtonActionPluggedIn{
		"hibernate":     Windows10GeneralConfigurationPowerButtonActionPluggedInhibernate,
		"noaction":      Windows10GeneralConfigurationPowerButtonActionPluggedInnoAction,
		"notconfigured": Windows10GeneralConfigurationPowerButtonActionPluggedInnotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerButtonActionPluggedInshutdown,
		"sleep":         Windows10GeneralConfigurationPowerButtonActionPluggedInsleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerButtonActionPluggedIn(input)
	return &out, nil
}
