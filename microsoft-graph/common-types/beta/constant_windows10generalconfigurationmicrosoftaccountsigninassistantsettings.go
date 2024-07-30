package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings string

const (
	Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings_Disabled      Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings = "disabled"
	Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings_NotConfigured Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings() []string {
	return []string{
		string(Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings_Disabled),
		string(Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings(input string) (*Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings, error) {
	vals := map[string]Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings{
		"disabled":      Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings_Disabled,
		"notconfigured": Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationMicrosoftAccountSignInAssistantSettings(input)
	return &out, nil
}
