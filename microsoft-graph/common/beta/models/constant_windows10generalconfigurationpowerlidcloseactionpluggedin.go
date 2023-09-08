package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerLidCloseActionPluggedIn string

const (
	Windows10GeneralConfigurationPowerLidCloseActionPluggedInhibernate     Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "Hibernate"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedInnoAction      Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "NoAction"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedInnotConfigured Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "NotConfigured"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedInshutdown      Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "Shutdown"
	Windows10GeneralConfigurationPowerLidCloseActionPluggedInsleep         Windows10GeneralConfigurationPowerLidCloseActionPluggedIn = "Sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerLidCloseActionPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedInhibernate),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedInnoAction),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedInnotConfigured),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedInshutdown),
		string(Windows10GeneralConfigurationPowerLidCloseActionPluggedInsleep),
	}
}

func parseWindows10GeneralConfigurationPowerLidCloseActionPluggedIn(input string) (*Windows10GeneralConfigurationPowerLidCloseActionPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerLidCloseActionPluggedIn{
		"hibernate":     Windows10GeneralConfigurationPowerLidCloseActionPluggedInhibernate,
		"noaction":      Windows10GeneralConfigurationPowerLidCloseActionPluggedInnoAction,
		"notconfigured": Windows10GeneralConfigurationPowerLidCloseActionPluggedInnotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerLidCloseActionPluggedInshutdown,
		"sleep":         Windows10GeneralConfigurationPowerLidCloseActionPluggedInsleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerLidCloseActionPluggedIn(input)
	return &out, nil
}
