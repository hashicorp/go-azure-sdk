package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationDeveloperUnlockSetting string

const (
	Windows10GeneralConfigurationDeveloperUnlockSetting_Allowed       Windows10GeneralConfigurationDeveloperUnlockSetting = "allowed"
	Windows10GeneralConfigurationDeveloperUnlockSetting_Blocked       Windows10GeneralConfigurationDeveloperUnlockSetting = "blocked"
	Windows10GeneralConfigurationDeveloperUnlockSetting_NotConfigured Windows10GeneralConfigurationDeveloperUnlockSetting = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationDeveloperUnlockSetting() []string {
	return []string{
		string(Windows10GeneralConfigurationDeveloperUnlockSetting_Allowed),
		string(Windows10GeneralConfigurationDeveloperUnlockSetting_Blocked),
		string(Windows10GeneralConfigurationDeveloperUnlockSetting_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationDeveloperUnlockSetting) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationDeveloperUnlockSetting(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationDeveloperUnlockSetting(input string) (*Windows10GeneralConfigurationDeveloperUnlockSetting, error) {
	vals := map[string]Windows10GeneralConfigurationDeveloperUnlockSetting{
		"allowed":       Windows10GeneralConfigurationDeveloperUnlockSetting_Allowed,
		"blocked":       Windows10GeneralConfigurationDeveloperUnlockSetting_Blocked,
		"notconfigured": Windows10GeneralConfigurationDeveloperUnlockSetting_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationDeveloperUnlockSetting(input)
	return &out, nil
}
