package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationLockScreenActivateAppsWithVoice string

const (
	Windows10GeneralConfigurationLockScreenActivateAppsWithVoicedisabled      Windows10GeneralConfigurationLockScreenActivateAppsWithVoice = "Disabled"
	Windows10GeneralConfigurationLockScreenActivateAppsWithVoiceenabled       Windows10GeneralConfigurationLockScreenActivateAppsWithVoice = "Enabled"
	Windows10GeneralConfigurationLockScreenActivateAppsWithVoicenotConfigured Windows10GeneralConfigurationLockScreenActivateAppsWithVoice = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationLockScreenActivateAppsWithVoice() []string {
	return []string{
		string(Windows10GeneralConfigurationLockScreenActivateAppsWithVoicedisabled),
		string(Windows10GeneralConfigurationLockScreenActivateAppsWithVoiceenabled),
		string(Windows10GeneralConfigurationLockScreenActivateAppsWithVoicenotConfigured),
	}
}

func parseWindows10GeneralConfigurationLockScreenActivateAppsWithVoice(input string) (*Windows10GeneralConfigurationLockScreenActivateAppsWithVoice, error) {
	vals := map[string]Windows10GeneralConfigurationLockScreenActivateAppsWithVoice{
		"disabled":      Windows10GeneralConfigurationLockScreenActivateAppsWithVoicedisabled,
		"enabled":       Windows10GeneralConfigurationLockScreenActivateAppsWithVoiceenabled,
		"notconfigured": Windows10GeneralConfigurationLockScreenActivateAppsWithVoicenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationLockScreenActivateAppsWithVoice(input)
	return &out, nil
}
