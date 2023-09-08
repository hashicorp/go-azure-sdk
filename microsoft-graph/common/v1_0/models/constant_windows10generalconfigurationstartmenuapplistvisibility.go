package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuAppListVisibility string

const (
	Windows10GeneralConfigurationStartMenuAppListVisibilitycollapse           Windows10GeneralConfigurationStartMenuAppListVisibility = "Collapse"
	Windows10GeneralConfigurationStartMenuAppListVisibilitydisableSettingsApp Windows10GeneralConfigurationStartMenuAppListVisibility = "DisableSettingsApp"
	Windows10GeneralConfigurationStartMenuAppListVisibilityremove             Windows10GeneralConfigurationStartMenuAppListVisibility = "Remove"
	Windows10GeneralConfigurationStartMenuAppListVisibilityuserDefined        Windows10GeneralConfigurationStartMenuAppListVisibility = "UserDefined"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuAppListVisibility() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuAppListVisibilitycollapse),
		string(Windows10GeneralConfigurationStartMenuAppListVisibilitydisableSettingsApp),
		string(Windows10GeneralConfigurationStartMenuAppListVisibilityremove),
		string(Windows10GeneralConfigurationStartMenuAppListVisibilityuserDefined),
	}
}

func parseWindows10GeneralConfigurationStartMenuAppListVisibility(input string) (*Windows10GeneralConfigurationStartMenuAppListVisibility, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuAppListVisibility{
		"collapse":           Windows10GeneralConfigurationStartMenuAppListVisibilitycollapse,
		"disablesettingsapp": Windows10GeneralConfigurationStartMenuAppListVisibilitydisableSettingsApp,
		"remove":             Windows10GeneralConfigurationStartMenuAppListVisibilityremove,
		"userdefined":        Windows10GeneralConfigurationStartMenuAppListVisibilityuserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuAppListVisibility(input)
	return &out, nil
}
