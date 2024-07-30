package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationLockScreenActivateAppsWithVoice string

const (
	Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_Disabled      Windows10GeneralConfigurationLockScreenActivateAppsWithVoice = "disabled"
	Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_Enabled       Windows10GeneralConfigurationLockScreenActivateAppsWithVoice = "enabled"
	Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_NotConfigured Windows10GeneralConfigurationLockScreenActivateAppsWithVoice = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationLockScreenActivateAppsWithVoice() []string {
	return []string{
		string(Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_Disabled),
		string(Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_Enabled),
		string(Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationLockScreenActivateAppsWithVoice) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationLockScreenActivateAppsWithVoice(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationLockScreenActivateAppsWithVoice(input string) (*Windows10GeneralConfigurationLockScreenActivateAppsWithVoice, error) {
	vals := map[string]Windows10GeneralConfigurationLockScreenActivateAppsWithVoice{
		"disabled":      Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_Disabled,
		"enabled":       Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_Enabled,
		"notconfigured": Windows10GeneralConfigurationLockScreenActivateAppsWithVoice_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationLockScreenActivateAppsWithVoice(input)
	return &out, nil
}
