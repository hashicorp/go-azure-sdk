package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows81GeneralConfigurationUserAccountControlSettings string

const (
	Windows81GeneralConfigurationUserAccountControlSettingsalwaysNotify                     Windows81GeneralConfigurationUserAccountControlSettings = "AlwaysNotify"
	Windows81GeneralConfigurationUserAccountControlSettingsneverNotify                      Windows81GeneralConfigurationUserAccountControlSettings = "NeverNotify"
	Windows81GeneralConfigurationUserAccountControlSettingsnotifyOnAppChanges               Windows81GeneralConfigurationUserAccountControlSettings = "NotifyOnAppChanges"
	Windows81GeneralConfigurationUserAccountControlSettingsnotifyOnAppChangesWithoutDimming Windows81GeneralConfigurationUserAccountControlSettings = "NotifyOnAppChangesWithoutDimming"
	Windows81GeneralConfigurationUserAccountControlSettingsuserDefined                      Windows81GeneralConfigurationUserAccountControlSettings = "UserDefined"
)

func PossibleValuesForWindows81GeneralConfigurationUserAccountControlSettings() []string {
	return []string{
		string(Windows81GeneralConfigurationUserAccountControlSettingsalwaysNotify),
		string(Windows81GeneralConfigurationUserAccountControlSettingsneverNotify),
		string(Windows81GeneralConfigurationUserAccountControlSettingsnotifyOnAppChanges),
		string(Windows81GeneralConfigurationUserAccountControlSettingsnotifyOnAppChangesWithoutDimming),
		string(Windows81GeneralConfigurationUserAccountControlSettingsuserDefined),
	}
}

func parseWindows81GeneralConfigurationUserAccountControlSettings(input string) (*Windows81GeneralConfigurationUserAccountControlSettings, error) {
	vals := map[string]Windows81GeneralConfigurationUserAccountControlSettings{
		"alwaysnotify":                     Windows81GeneralConfigurationUserAccountControlSettingsalwaysNotify,
		"nevernotify":                      Windows81GeneralConfigurationUserAccountControlSettingsneverNotify,
		"notifyonappchanges":               Windows81GeneralConfigurationUserAccountControlSettingsnotifyOnAppChanges,
		"notifyonappchangeswithoutdimming": Windows81GeneralConfigurationUserAccountControlSettingsnotifyOnAppChangesWithoutDimming,
		"userdefined":                      Windows81GeneralConfigurationUserAccountControlSettingsuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows81GeneralConfigurationUserAccountControlSettings(input)
	return &out, nil
}
