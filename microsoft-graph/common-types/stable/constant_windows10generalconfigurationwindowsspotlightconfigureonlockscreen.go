package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen string

const (
	Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_Disabled      Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen = "disabled"
	Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_Enabled       Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen = "enabled"
	Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_NotConfigured Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen() []string {
	return []string{
		string(Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_Disabled),
		string(Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_Enabled),
		string(Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen(input string) (*Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen, error) {
	vals := map[string]Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen{
		"disabled":      Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_Disabled,
		"enabled":       Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_Enabled,
		"notconfigured": Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationWindowsSpotlightConfigureOnLockScreen(input)
	return &out, nil
}
