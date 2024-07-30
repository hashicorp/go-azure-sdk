package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationActivateAppsWithVoice string

const (
	Windows10GeneralConfigurationActivateAppsWithVoice_Disabled      Windows10GeneralConfigurationActivateAppsWithVoice = "disabled"
	Windows10GeneralConfigurationActivateAppsWithVoice_Enabled       Windows10GeneralConfigurationActivateAppsWithVoice = "enabled"
	Windows10GeneralConfigurationActivateAppsWithVoice_NotConfigured Windows10GeneralConfigurationActivateAppsWithVoice = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationActivateAppsWithVoice() []string {
	return []string{
		string(Windows10GeneralConfigurationActivateAppsWithVoice_Disabled),
		string(Windows10GeneralConfigurationActivateAppsWithVoice_Enabled),
		string(Windows10GeneralConfigurationActivateAppsWithVoice_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationActivateAppsWithVoice) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationActivateAppsWithVoice(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationActivateAppsWithVoice(input string) (*Windows10GeneralConfigurationActivateAppsWithVoice, error) {
	vals := map[string]Windows10GeneralConfigurationActivateAppsWithVoice{
		"disabled":      Windows10GeneralConfigurationActivateAppsWithVoice_Disabled,
		"enabled":       Windows10GeneralConfigurationActivateAppsWithVoice_Enabled,
		"notconfigured": Windows10GeneralConfigurationActivateAppsWithVoice_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationActivateAppsWithVoice(input)
	return &out, nil
}
