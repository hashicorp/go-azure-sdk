package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerHybridSleepPluggedIn string

const (
	Windows10GeneralConfigurationPowerHybridSleepPluggedIn_Disabled      Windows10GeneralConfigurationPowerHybridSleepPluggedIn = "disabled"
	Windows10GeneralConfigurationPowerHybridSleepPluggedIn_Enabled       Windows10GeneralConfigurationPowerHybridSleepPluggedIn = "enabled"
	Windows10GeneralConfigurationPowerHybridSleepPluggedIn_NotConfigured Windows10GeneralConfigurationPowerHybridSleepPluggedIn = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationPowerHybridSleepPluggedIn() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerHybridSleepPluggedIn_Disabled),
		string(Windows10GeneralConfigurationPowerHybridSleepPluggedIn_Enabled),
		string(Windows10GeneralConfigurationPowerHybridSleepPluggedIn_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationPowerHybridSleepPluggedIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerHybridSleepPluggedIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerHybridSleepPluggedIn(input string) (*Windows10GeneralConfigurationPowerHybridSleepPluggedIn, error) {
	vals := map[string]Windows10GeneralConfigurationPowerHybridSleepPluggedIn{
		"disabled":      Windows10GeneralConfigurationPowerHybridSleepPluggedIn_Disabled,
		"enabled":       Windows10GeneralConfigurationPowerHybridSleepPluggedIn_Enabled,
		"notconfigured": Windows10GeneralConfigurationPowerHybridSleepPluggedIn_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerHybridSleepPluggedIn(input)
	return &out, nil
}
