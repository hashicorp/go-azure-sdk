package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10GeneralConfigurationPowerSleepButtonActionOnBattery string

const (
	Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Hibernate     Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "hibernate"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_NoAction      Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "noAction"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_NotConfigured Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "notConfigured"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Shutdown      Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "shutdown"
	Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Sleep         Windows10GeneralConfigurationPowerSleepButtonActionOnBattery = "sleep"
)

func PossibleValuesForWindows10GeneralConfigurationPowerSleepButtonActionOnBattery() []string {
	return []string{
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Hibernate),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_NoAction),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_NotConfigured),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Shutdown),
		string(Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Sleep),
	}
}

func (s *Windows10GeneralConfigurationPowerSleepButtonActionOnBattery) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10GeneralConfigurationPowerSleepButtonActionOnBattery(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10GeneralConfigurationPowerSleepButtonActionOnBattery(input string) (*Windows10GeneralConfigurationPowerSleepButtonActionOnBattery, error) {
	vals := map[string]Windows10GeneralConfigurationPowerSleepButtonActionOnBattery{
		"hibernate":     Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Hibernate,
		"noaction":      Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_NoAction,
		"notconfigured": Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_NotConfigured,
		"shutdown":      Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Shutdown,
		"sleep":         Windows10GeneralConfigurationPowerSleepButtonActionOnBattery_Sleep,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10GeneralConfigurationPowerSleepButtonActionOnBattery(input)
	return &out, nil
}
