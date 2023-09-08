package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting string

const (
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingauditMode     Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "AuditMode"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingenable        Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "Enable"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingnotConfigured Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "NotConfigured"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettinguserDefined   Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "UserDefined"
	Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingwarn          Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting = "Warn"
)

func PossibleValuesForWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting() []string {
	return []string{
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingauditMode),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingenable),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingnotConfigured),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettinguserDefined),
		string(Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingwarn),
	}
}

func parseWindows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting(input string) (*Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting, error) {
	vals := map[string]Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting{
		"auditmode":     Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingauditMode,
		"enable":        Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingenable,
		"notconfigured": Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingnotConfigured,
		"userdefined":   Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettinguserDefined,
		"warn":          Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSettingwarn,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDefenderPotentiallyUnwantedAppActionSetting(input)
	return &out, nil
}
