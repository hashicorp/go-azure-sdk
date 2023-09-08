package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationStartMenuPinnedFolderSettings string

const (
	Windows10GeneralConfigurationStartMenuPinnedFolderSettingshide          Windows10GeneralConfigurationStartMenuPinnedFolderSettings = "Hide"
	Windows10GeneralConfigurationStartMenuPinnedFolderSettingsnotConfigured Windows10GeneralConfigurationStartMenuPinnedFolderSettings = "NotConfigured"
	Windows10GeneralConfigurationStartMenuPinnedFolderSettingsshow          Windows10GeneralConfigurationStartMenuPinnedFolderSettings = "Show"
)

func PossibleValuesForWindows10GeneralConfigurationStartMenuPinnedFolderSettings() []string {
	return []string{
		string(Windows10GeneralConfigurationStartMenuPinnedFolderSettingshide),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderSettingsnotConfigured),
		string(Windows10GeneralConfigurationStartMenuPinnedFolderSettingsshow),
	}
}

func parseWindows10GeneralConfigurationStartMenuPinnedFolderSettings(input string) (*Windows10GeneralConfigurationStartMenuPinnedFolderSettings, error) {
	vals := map[string]Windows10GeneralConfigurationStartMenuPinnedFolderSettings{
		"hide":          Windows10GeneralConfigurationStartMenuPinnedFolderSettingshide,
		"notconfigured": Windows10GeneralConfigurationStartMenuPinnedFolderSettingsnotConfigured,
		"show":          Windows10GeneralConfigurationStartMenuPinnedFolderSettingsshow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationStartMenuPinnedFolderSettings(input)
	return &out, nil
}
