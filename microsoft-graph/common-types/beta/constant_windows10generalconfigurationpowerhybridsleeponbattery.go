package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerHybridSleepOnBattery string

const (
	Windows10GeneralConfigurationPowerHybridSleepOnBattery_Disabled      Windows10GeneralConfigurationPowerHybridSleepOnBattery = "disabled"
	Windows10GeneralConfigurationPowerHybridSleepOnBattery_Enabled       Windows10GeneralConfigurationPowerHybridSleepOnBattery = "enabled"
	Windows10GeneralConfigurationPowerHybridSleepOnBattery_NotConfigured Windows10GeneralConfigurationPowerHybridSleepOnBattery = "notConfigured"
)

func PossibleValuesForWindows10GeneralConfigurationPowerHybridSleepOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerHybridSleepOnBattery_Disabled),
		string(Windows10GeneralConfigurationPowerHybridSleepOnBattery_Enabled),
		string(Windows10GeneralConfigurationPowerHybridSleepOnBattery_NotConfigured),
	}
}

func (s *Windows10GeneralConfigurationPowerHybridSleepOnBattery) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerHybridSleepOnBattery(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerHybridSleepOnBattery(input string) (*Windows10GeneralConfigurationPowerHybridSleepOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerHybridSleepOnBattery{
		"disabled":      Windows10GeneralConfigurationPowerHybridSleepOnBattery_Disabled,
		"enabled":       Windows10GeneralConfigurationPowerHybridSleepOnBattery_Enabled,
		"notconfigured": Windows10GeneralConfigurationPowerHybridSleepOnBattery_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerHybridSleepOnBattery(input)
	return &out, nil
}
