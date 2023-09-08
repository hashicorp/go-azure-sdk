package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDeveloperUnlockSetting string

const (
	Windows10GeneralConfigurationDeveloperUnlockSettingallowed       Windows10GeneralConfigurationDeveloperUnlockSetting = "Allowed"
	Windows10GeneralConfigurationDeveloperUnlockSettingblocked       Windows10GeneralConfigurationDeveloperUnlockSetting = "Blocked"
	Windows10GeneralConfigurationDeveloperUnlockSettingnotConfigured Windows10GeneralConfigurationDeveloperUnlockSetting = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationDeveloperUnlockSetting() []string {
	return []string{
		string(Windows10GeneralConfigurationDeveloperUnlockSettingallowed),
		string(Windows10GeneralConfigurationDeveloperUnlockSettingblocked),
		string(Windows10GeneralConfigurationDeveloperUnlockSettingnotConfigured),
	}
}

func parseWindows10GeneralConfigurationDeveloperUnlockSetting(input string) (*Windows10GeneralConfigurationDeveloperUnlockSetting, error) {
	vals := map[string]Windows10GeneralConfigurationDeveloperUnlockSetting{
		"allowed":       Windows10GeneralConfigurationDeveloperUnlockSettingallowed,
		"blocked":       Windows10GeneralConfigurationDeveloperUnlockSettingblocked,
		"notconfigured": Windows10GeneralConfigurationDeveloperUnlockSettingnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDeveloperUnlockSetting(input)
	return &out, nil
}
