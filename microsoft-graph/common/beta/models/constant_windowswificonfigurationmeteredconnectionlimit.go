package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiConfigurationMeteredConnectionLimit string

const (
	WindowsWifiConfigurationMeteredConnectionLimitfixed        WindowsWifiConfigurationMeteredConnectionLimit = "Fixed"
	WindowsWifiConfigurationMeteredConnectionLimitunrestricted WindowsWifiConfigurationMeteredConnectionLimit = "Unrestricted"
	WindowsWifiConfigurationMeteredConnectionLimitvariable     WindowsWifiConfigurationMeteredConnectionLimit = "Variable"
)

func PossibleValuesForWindowsWifiConfigurationMeteredConnectionLimit() []string {
	return []string{
		string(WindowsWifiConfigurationMeteredConnectionLimitfixed),
		string(WindowsWifiConfigurationMeteredConnectionLimitunrestricted),
		string(WindowsWifiConfigurationMeteredConnectionLimitvariable),
	}
}

func parseWindowsWifiConfigurationMeteredConnectionLimit(input string) (*WindowsWifiConfigurationMeteredConnectionLimit, error) {
	vals := map[string]WindowsWifiConfigurationMeteredConnectionLimit{
		"fixed":        WindowsWifiConfigurationMeteredConnectionLimitfixed,
		"unrestricted": WindowsWifiConfigurationMeteredConnectionLimitunrestricted,
		"variable":     WindowsWifiConfigurationMeteredConnectionLimitvariable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiConfigurationMeteredConnectionLimit(input)
	return &out, nil
}
