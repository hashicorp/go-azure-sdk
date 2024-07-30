package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerButtonActionOnBattery string

const (
	Windows10GeneralConfigurationPowerButtonActionOnBattery_Hibernate     Windows10GeneralConfigurationPowerButtonActionOnBattery = "hibernate"
	Windows10GeneralConfigurationPowerButtonActionOnBattery_NoAction      Windows10GeneralConfigurationPowerButtonActionOnBattery = "noAction"
	Windows10GeneralConfigurationPowerButtonActionOnBattery_NotConfigured Windows10GeneralConfigurationPowerButtonActionOnBattery = "notConfigured"
	Windows10GeneralConfigurationPowerButtonActionOnBattery_Shutdown      Windows10GeneralConfigurationPowerButtonActionOnBattery = "shutdown"
	Windows10GeneralConfigurationPowerButtonActionOnBattery_Sleep         Windows10GeneralConfigurationPowerButtonActionOnBattery = "sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerButtonActionOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerButtonActionOnBattery_Hibernate),
		string(Windows10GeneralConfigurationPowerButtonActionOnBattery_NoAction),
		string(Windows10GeneralConfigurationPowerButtonActionOnBattery_NotConfigured),
		string(Windows10GeneralConfigurationPowerButtonActionOnBattery_Shutdown),
		string(Windows10GeneralConfigurationPowerButtonActionOnBattery_Sleep),
	}
}

func (s *Windows10GeneralConfigurationPowerButtonActionOnBattery) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerButtonActionOnBattery(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerButtonActionOnBattery(input string) (*Windows10GeneralConfigurationPowerButtonActionOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerButtonActionOnBattery{
		"hibernate":     Windows10GeneralConfigurationPowerButtonActionOnBattery_Hibernate,
		"noaction":      Windows10GeneralConfigurationPowerButtonActionOnBattery_NoAction,
		"notconfigured": Windows10GeneralConfigurationPowerButtonActionOnBattery_NotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerButtonActionOnBattery_Shutdown,
		"sleep":         Windows10GeneralConfigurationPowerButtonActionOnBattery_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerButtonActionOnBattery(input)
	return &out, nil
}
