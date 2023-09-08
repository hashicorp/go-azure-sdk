package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings string

const (
	Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsblocked                 Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings = "Blocked"
	Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsblockedWithUserOverride Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings = "BlockedWithUserOverride"
	Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsnotConfigured           Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationExperienceDoNotSyncBrowserSettings() []string {
	return []string{
		string(Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsblocked),
		string(Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsblockedWithUserOverride),
		string(Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsnotConfigured),
	}
}

func parseWindows10GeneralConfigurationExperienceDoNotSyncBrowserSettings(input string) (*Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings, error) {
	vals := map[string]Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings{
		"blocked":                 Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsblocked,
		"blockedwithuseroverride": Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsblockedWithUserOverride,
		"notconfigured":           Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettingsnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings(input)
	return &out, nil
}
