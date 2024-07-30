package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings string

const (
	Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_Blocked                 Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings = "blocked"
	Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_BlockedWithUserOverride Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings = "blockedWithUserOverride"
	Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_NotConfigured           Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationExperienceDoNotSyncBrowserSettings() []string {
	return []string{
		string(Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_Blocked),
		string(Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_BlockedWithUserOverride),
		string(Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationExperienceDoNotSyncBrowserSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationExperienceDoNotSyncBrowserSettings(input string) (*Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings, error) {
	vals := map[string]Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings{
		"blocked":                 Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_Blocked,
		"blockedwithuseroverride": Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_BlockedWithUserOverride,
		"notconfigured":           Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationExperienceDoNotSyncBrowserSettings(input)
	return &out, nil
}
