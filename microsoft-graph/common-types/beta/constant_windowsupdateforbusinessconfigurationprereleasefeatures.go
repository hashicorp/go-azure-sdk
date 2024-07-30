package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfigurationPrereleaseFeatures string

const (
	WindowsUpdateForBusinessConfigurationPrereleaseFeatures_NotAllowed                  WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "notAllowed"
	WindowsUpdateForBusinessConfigurationPrereleaseFeatures_SettingsAndExperimentations WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "settingsAndExperimentations"
	WindowsUpdateForBusinessConfigurationPrereleaseFeatures_SettingsOnly                WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "settingsOnly"
	WindowsUpdateForBusinessConfigurationPrereleaseFeatures_UserDefined                 WindowsUpdateForBusinessConfigurationPrereleaseFeatures = "userDefined"
)

func PossibleValuesForWindowsUpdateForBusinessConfigurationPrereleaseFeatures() []string {
	return []string{
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeatures_NotAllowed),
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeatures_SettingsAndExperimentations),
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeatures_SettingsOnly),
		string(WindowsUpdateForBusinessConfigurationPrereleaseFeatures_UserDefined),
	}
}

func (s *WindowsUpdateForBusinessConfigurationPrereleaseFeatures) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsUpdateForBusinessConfigurationPrereleaseFeatures(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsUpdateForBusinessConfigurationPrereleaseFeatures(input string) (*WindowsUpdateForBusinessConfigurationPrereleaseFeatures, error) {
	vals := map[string]WindowsUpdateForBusinessConfigurationPrereleaseFeatures{
		"notallowed":                  WindowsUpdateForBusinessConfigurationPrereleaseFeatures_NotAllowed,
		"settingsandexperimentations": WindowsUpdateForBusinessConfigurationPrereleaseFeatures_SettingsAndExperimentations,
		"settingsonly":                WindowsUpdateForBusinessConfigurationPrereleaseFeatures_SettingsOnly,
		"userdefined":                 WindowsUpdateForBusinessConfigurationPrereleaseFeatures_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsUpdateForBusinessConfigurationPrereleaseFeatures(input)
	return &out, nil
}
