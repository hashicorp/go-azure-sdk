package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationActivateAppsWithVoice string

const (
	Windows10GeneralConfigurationActivateAppsWithVoicedisabled      Windows10GeneralConfigurationActivateAppsWithVoice = "Disabled"
	Windows10GeneralConfigurationActivateAppsWithVoiceenabled       Windows10GeneralConfigurationActivateAppsWithVoice = "Enabled"
	Windows10GeneralConfigurationActivateAppsWithVoicenotConfigured Windows10GeneralConfigurationActivateAppsWithVoice = "NotConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationActivateAppsWithVoice() []string {
	return []string{
		string(Windows10GeneralConfigurationActivateAppsWithVoicedisabled),
		string(Windows10GeneralConfigurationActivateAppsWithVoiceenabled),
		string(Windows10GeneralConfigurationActivateAppsWithVoicenotConfigured),
	}
}

func parseWindows10GeneralConfigurationActivateAppsWithVoice(input string) (*Windows10GeneralConfigurationActivateAppsWithVoice, error) {
	vals := map[string]Windows10GeneralConfigurationActivateAppsWithVoice{
		"disabled":      Windows10GeneralConfigurationActivateAppsWithVoicedisabled,
		"enabled":       Windows10GeneralConfigurationActivateAppsWithVoiceenabled,
		"notconfigured": Windows10GeneralConfigurationActivateAppsWithVoicenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationActivateAppsWithVoice(input)
	return &out, nil
}
